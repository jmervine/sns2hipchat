package sns

import (
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	. "github.com/jmervine/sns2hipchat/Godeps/_workspace/src/gopkg.in/jmervine/GoT.v1"
)

const STUB_FILE = "_support/notification.json"

var note Notification

func stubReader() io.ReadCloser {
	file, err := os.Open("_support/notification.json")
	if err != nil {
		panic(err) // should never happen
	}

	return io.ReadCloser(file)
}

func TestParseRequestBody(T *testing.T) {
	_, err := ParseRequestBody(nil)
	Go(T).RefuteNil(err)

	reader := stubReader()
	note, err = ParseRequestBody(reader)

	Go(T).AssertNil(err)
	Go(T).RefuteNil(note)
	Go(T).AssertEqual(note.Message, "Message")
}

func TestNotification_ToString(T *testing.T) {
	reader := stubReader()
	note, err := ParseRequestBody(reader)

	Go(T).AssertNil(err)
	Go(T).RefuteNil(note)
	Go(T).AssertEqual(note.ToString(), "Subject: Message")

	note.Subject = ""
	note.Message = ""
	Go(T).AssertEqual(note.ToString(), "<no subject>: <no message>")
}

func TestNotification_ToJson(T *testing.T) {
	reader := stubReader()
	note, err := ParseRequestBody(reader)

	Go(T).AssertNil(err)
	Go(T).RefuteNil(note)

	s, err := note.ToJson()

	Go(T).AssertNil(err)

	n := Notification{}
	dec := json.NewDecoder(strings.NewReader(s))
	err = dec.Decode(&n)

	Go(T).AssertNil(err)
	Go(T).AssertEqual(n, note)
}
