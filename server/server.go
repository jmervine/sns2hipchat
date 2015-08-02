package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/jmervine/sns2hipchat/config"
	"github.com/jmervine/sns2hipchat/formatters"
	"github.com/jmervine/sns2hipchat/hipchat"
	"github.com/jmervine/sns2hipchat/sns"
)

var Formatter formatter.Formatter

func Start(cfg *config.Config) error {
	s := &http.Server{
		Addr: cfg.Addr,
	}

	// TODO: I should be able use interfaces for this somehow.
	var err error
	var v1 *hipchat.HipchatV1
	var v2 *hipchat.HipchatV2

	if cfg.HipchatAPIVersion == 1 {
		v1, err = hipchat.NewV1(cfg)
	} else if cfg.HipchatAPIVersion == 2 {
		v2, err = hipchat.NewV2(cfg)
	} else {
		err = fmt.Errorf("unhandled hipchat api version")
	}

	if err != nil {
		panic(err)
	}

	s.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()

		// defer logging
		defer func() {
			if !cfg.Debug { // allow panics in debug mode
				if e := recover(); e != nil {
					http.Error(w, e.(error).Error(), http.StatusInternalServerError)
					fmt.Printf("at=request status=500 method=%s url=%s error=%s took=%v\n",
						r.Method,
						r.URL.Path,
						e.(error).Error(),
						time.Since(begin))
				} else {
					fmt.Printf("at=request status=200 method=%s url=%s took=%v\n",
						r.Method,
						r.URL.Path,
						time.Since(begin))
				}
			} else {
				fmt.Printf("DEBUG:: at=request request=%+v\n", *r)
			}
		}()

		var rooms []string
		var err error
		var notification sns.Notification

		if cfg.Test && r.URL.Path == "/test" {
			f := "sns/_support/message.json"
			var file io.ReadCloser
			file, err = os.Open(f)
			if err != nil { //&& err != io.EOF {
				panic(err) // should never happen
			}

			notification, err = sns.ParseRequestBody(file)
		} else if r.Method == "POST" { // ensure post
			notification, err = sns.ParseRequestBody(r.Body)
		} else {
			http.NotFound(w, r)
			fmt.Printf("at=request status=404 method=%s url=%s took=%v\n",
				r.Method,
				r.URL.Path,
				time.Since(begin))
			return
		}

		if err != nil {
			panic(err) // trigger error logging
		}

		if _, ok := r.URL.Query()["room"]; ok {
			rooms = r.URL.Query()["room"]
		}

		if len(rooms) == 0 {
			rooms = cfg.Rooms
		}

		if cfg.Debug {
			fmt.Printf("DEBUG:: at=request notification=%+v\n", notification)
		}

		if notification.HandleSubURL() {
			return
		}

		// TODO:
		// I feel like formatting should be handled either in hipchat/hipchat.go
		// or sns/notification.go, not here.
		if Formatter == nil {
			Formatter = formatter.New(cfg.Formatter)
		}

		var m string
		if cfg.MessageFormat == "html" {
			m, err = Formatter.FormatHTML(&notification)
		} else {
			m, err = Formatter.Format(&notification)
		}

		if err != nil {
			panic(err) // trigger error logging
		}

		if cfg.Test || cfg.Debug {
			fmt.Printf("DEBUG:: at=request target=hipchat message=%s\n", m)
			if cfg.Test {
				fmt.Fprintf(w, "<html><body>%s</body></html>", m)
				return
			}
		}

		if v1 != nil {
			err = v1.Post(rooms, m)
		} else {
			err = v2.Post(rooms, m)
		}

		if err != nil {
			panic(err) // trigger error logging
		}
	})

	if cfg.Debug {
		fmt.Printf("at=startup config=%+v\n", *cfg)
	} else {
		fmt.Printf("at=startup listen=%v hipchat=%v\n", cfg.Addr, cfg.Host)
	}
	return s.ListenAndServe()
}
