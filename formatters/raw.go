package formatter

import (
	"fmt"

	"github.com/jmervine/sns2hipchat/sns"
)

type Raw struct {
	Formatter
}

func (f Raw) Format(n *sns.Notification) (msg string, err error) {
	return n.ToJson()
}

func (f Raw) FormatHTML(n *sns.Notification) (msg string, err error) {
	msg, err = f.Format(n)
	msg = fmt.Sprintf("<pre>%s\n</pre>", msg)

	return
}
