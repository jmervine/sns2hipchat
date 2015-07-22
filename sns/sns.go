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
	dec := json.NewDecoder(body)
	err = dec.Decode(&n)

	return
}

func (n Notification) ToString() string {
	var sub, msg string
	if n.Subject == "" {
		sub = "<no subject>"
	}

	if n.Message == "" {
		msg = "<no message>"
	}

	return fmt.Sprintf("%s: %s",
		sub,
		msg,
	)
}
