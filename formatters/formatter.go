package formatter

import (
	"fmt"
	"github.com/jmervine/sns2hipchat/sns"
)

type Formatter interface {
	FormatHTML(n *sns.Notification) (string, error)
	Format(n *sns.Notification) (string, error)
}

// Return formatter interface, currently supported:
//
// * basic - Basic formatter
// * alarm - CloudWatchAlarm formatter
func New(name string) (f Formatter) {
	// switch to handle many
	switch name {
	case "basic":
		f = new(Basic)
	case "alarm":
		f = new(CloudWatchAlarm)
	}

	if f == nil {
		panic(fmt.Errorf("unhandled formatter: %v", name))
	}

	return
}

type Basic struct {
	Formatter
}

func (f Basic) Format(n *sns.Notification) (msg string, err error) {
	s := n.Subject
	m := n.Message
	if s == "" {
		s = "<no subject>"
	}

	if m == "" {
		m = "<no message>"
	}

	msg = fmt.Sprintf("%s: %s", s, m)

	return
}

func (f Basic) FormatHTML(n *sns.Notification) (msg string, err error) {
	s := n.Subject
	m := n.Message

	msg = fmt.Sprintf(`<b>Subject:</b>: %s
<hr />
<b>Message:</b>
<br />
%s
<br />`, s, m)

	return
}
