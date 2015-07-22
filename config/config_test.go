package config

import (
	"testing"

	. "github.com/jmervine/hipchat-sns-relay/Godeps/_workspace/src/gopkg.in/jmervine/GoT.v1"
)

var cfg *Config

var required = []string{"app", "-t", "token", "-r", "room"}

func TestParseRequired(T *testing.T) {
	cfg = Parse([]string{"app"})
	Go(T).AssertNil(cfg)

	cfg = Parse([]string{"app", "-t", "token"})
	Go(T).AssertNil(cfg)

	cfg = Parse([]string{"app", "-r", "room"})
	Go(T).AssertNil(cfg)

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
