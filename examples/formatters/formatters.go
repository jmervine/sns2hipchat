package main

import (
	"fmt"
	"os"

	// Import all the things...
	"github.com/jmervine/sns2hipchat/config"
	"github.com/jmervine/sns2hipchat/formatters"
	"github.com/jmervine/sns2hipchat/server"
	"github.com/jmervine/sns2hipchat/sns"
)

// HOW TO: Creating and using a custom formatter.

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
	// Parse args to configuration.
	if cfg := config.Parse(os.Args); cfg != nil {

		// Start server with configuration.
		server.Start(cfg)
	}
}
