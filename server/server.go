package server

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/jmervine/sns2hipchat/config"
	"github.com/jmervine/sns2hipchat/hipchat"
	"github.com/jmervine/sns2hipchat/sns"
)

func Start(cfg *config.Config) error {
	if cfg.Debug {
		fmt.Printf("at=startup config=%+v\n", *cfg)
	}

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
			if e := recover(); e != nil {
				http.Error(w, e.(error).Error(), http.StatusInternalServerError)
				fmt.Printf("at=request status=500 error=%s took=%v\n",
					e.(error).Error(),
					time.Since(begin))
				if cfg.Debug {
					fmt.Printf("DEBUG::\n%s: %s\n\n", e, debug.Stack())
				}
			} else {
				fmt.Printf("at=request status=200 url=%s took=%v\n",
					r.URL.Path,
					time.Since(begin))
			}

			if cfg.Debug {
				fmt.Printf("DEBUG:: at=request request=%+v\n", *r)
			}
		}()

		notification, err := sns.ParseRequestBody(r.Body)

		if err != nil {
			panic(err) // trigger error logging
		}

		if cfg.Debug {
			fmt.Printf("DEBUG:: at=request notification=%+v\n", notification)
		}

		if notification.HandleSubURL() {
			return
		}

		m := notification.ToString()
		if v1 != nil {
			err = v1.Post(cfg.RoomID, m)
		} else {
			err = v2.Post(cfg.RoomID, m)
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
