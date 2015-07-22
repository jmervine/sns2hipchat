package main

import (
	"github.com/jmervine/hipchat-sns-relay/config"
	"github.com/jmervine/hipchat-sns-relay/server"
	"os"
)

func main() {
	server.Start(config.Parse(os.Args))
}
