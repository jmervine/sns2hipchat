package main

import (
	"github.com/jmervine/sns2hipchat/config"
	"github.com/jmervine/sns2hipchat/server"
	"os"
)

func main() {
	cfg := config.Parse(os.Args)

	if cfg != nil {
		server.Start(cfg)
	}
}
