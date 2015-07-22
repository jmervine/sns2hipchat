package main

import (
	"github.com/jmervine/hipchat-sns-relay/config"
	"github.com/jmervine/hipchat-sns-relay/server"
	"os"
)

func main() {
	cfg := config.Parse(os.Args)

	if cfg != nil {
		server.Start(cfg)
	}
}
