package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmervine/hipchat-sns-relay/config"
	"github.com/jmervine/hipchat-sns-relay/hipchat"
	"github.com/jmervine/hipchat-sns-relay/sns"
)

func Start(cfg *config.Config) error {
	s := &http.Server{
		Addr: cfg.Addr,
	}

	s.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()

		// defer logging
		defer func() {
			if e := recover(); e != nil {
				http.Error(w, e.(error).Error(), http.StatusInternalServerError)
				fmt.Printf("at=request status=500 error=%s duration=%v\n",
					e.(error).Error(),
					time.Since(begin))
			} else {
				fmt.Printf("at=request status=200 url=%s\n",
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

		h, err := hipchat.New(cfg)
		if err != nil {
			panic(err) // trigger error logging
		}

		err = h.Post(notification.ToString())
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
