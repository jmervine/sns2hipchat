package formatter

import (
	"fmt"
	"github.com/jmervine/sns2hipchat/sns"
)

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

	msg = fmt.Sprintf(`<b>Subject:</b> %s
<br />
<b>Message:</b>
<br />
%s
<br />`, s, m)

	return
}
