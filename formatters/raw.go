package formatter

import (
	"encoding/json"
	"fmt"

	"github.com/jmervine/sns2hipchat/sns"
)

type Raw struct {
	Formatter
}

func (f Raw) Format(n *sns.Notification) (msg string, err error) {
	b, err := json.MarshalIndent(n, "", "  ")

	msg = string(b)

	return
}

func (f Raw) FormatHTML(n *sns.Notification) (msg string, err error) {
	msg, err = f.Format(n)
	msg = fmt.Sprintf("<pre><code>%s\n</code></pre>", msg)

	return
}
