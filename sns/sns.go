package sns

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Notification struct {
	Message          string
	MessageId        string
	Signature        string
	SignatureVersion string
	SigningCertURL   string
	SubscribeURL     string
	Subject          string
	Timestamp        string
	TopicArn         string
	Type             string
	UnsubscribeURL   string
}

func ParseRequestBody(body io.ReadCloser) (n Notification, err error) {
	defer func() {
		// don't panic
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	dec := json.NewDecoder(body)
	err = dec.Decode(&n)

	return
}

func (n Notification) HandleSubURL() bool {
	if s := n.SubscribeURL; len(s) != 0 {
		fmt.Printf("SubscribeURL detected: %v\n", s)

		if _, err := http.Get(s); err != nil {
			fmt.Printf("Subscribe error: %v\n", err)
		}
		return true
	}
	return false
}

func (n Notification) ToString() string {
	sub := n.Subject
	msg := n.Message
	if sub == "" {
		sub = "<no subject>"
	}

	if msg == "" {
		msg = "<no message>"
	}

	return fmt.Sprintf("%s: %s",
		sub,
		msg,
	)
}

func (n *Notification) ToJson() (string, error) {
	b, err := json.MarshalIndent(n, "", "  ")

	msg := string(b)
	return msg, err
}
