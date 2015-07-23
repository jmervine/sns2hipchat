package main

/**
 * HOW-TO: Creating a custom formatter:
 *
 * To add a custom formatter, you create your own application, using
 * sns2hipchat internals and build a formatter struct using the Formatter
 * interface, as shown here.
 **/

import (
	"fmt"
	"os"

	// Import all the things...
	//
	// Import configuration, environment handling and cli.
	"github.com/jmervine/sns2hipchat/config"

	// Import formatters interface.
	"github.com/jmervine/sns2hipchat/formatters"

	// Import http server handling.
	"github.com/jmervine/sns2hipchat/server"

	// Import SNS message handling.
	"github.com/jmervine/sns2hipchat/sns"
)

// Define Custom struct using the Formatter interface, exactly like so:
type Custom struct {
	formatter.Formatter
}

// Define Format func, which is a required method and will be chosen if
// '$HIPCHAT_FORMAT' is 'text'.
func (f Custom) Format(n *sns.Notification) (msg string, err error) {
	msg = fmt.Sprintf("Custom Formatter\nSubject: %s\nMessage: %s", n.Subject, n.Message)
	return
}

// Define FormatHTML func, which is a required method and will be chosen if
// '$HIPCHAT_FORMAT' is 'html'.
func (f Custom) FormatHTML(n *sns.Notification) (msg string, err error) {
	msg, err = f.Format(n)
	msg = fmt.Sprintf("<pre>%s</pre>", msg)
	return
}

func main() {
	// Set `server.Formatter` to your Custom formatter.
	server.Formatter = new(Custom)

	// This is from sns2hipchat's main.go...
	//
	// Parse args and environment vars to configuration.
	if cfg := config.Parse(os.Args); cfg != nil {

		// Start server with configuration.
		server.Start(cfg)
	}
}
