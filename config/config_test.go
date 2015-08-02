package config

import (
	"bytes"
	"io"
	"os"
	"testing"

	. "github.com/jmervine/sns2hipchat/Godeps/_workspace/src/gopkg.in/jmervine/GoT.v1"
)

var cfg *Config

var required = []string{"app", "-t", "token", "-r", "room"}

func TestParseRequired(T *testing.T) {
	// capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// call that prints to stdout
	cfg = Parse([]string{"app"})
	Go(T).AssertNil(cfg)

	// restore stdout
	w.Close()
	os.Stdout = old

	// capture stdout as string
	out := <-outC

	Go(T).AssertEqual(out, "Hipchat Token Required. See '--help' for details.")

	cfg = Parse(required)
	Go(T).RefuteNil(cfg)

}

func TestParseNotify(T *testing.T) {
	rcp := *(&required) // copy required
	rcp = append(rcp, "-n", "bad_value")

	cfg = Parse(rcp)
	Go(T).AssertNil(cfg)

	rcp = *(&required) // copy required
	rcp = append(rcp, "-n", "false")

	cfg = Parse(rcp)
	Go(T).RefuteNil(cfg)
	Go(T).Refute(cfg.Notify)

	rcp = *(&required) // copy required
	rcp = append(rcp, "-n", "true")

	cfg = Parse(rcp)
	Go(T).RefuteNil(cfg)
	Go(T).Assert(cfg.Notify)
}
