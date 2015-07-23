package main

import (
	"github.com/jmervine/sns2hipchat/config"
	"github.com/jmervine/sns2hipchat/server"
	"os"
)

func main() {
	// Parse args to configuration.
	if cfg := config.Parse(os.Args); cfg != nil {

		// Start server with configuration.
		server.Start(cfg)
	}
}
