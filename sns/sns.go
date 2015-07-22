package sns

import (
	"encoding/json"
	"fmt"
	"io"
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
