package hipchat

import (
	"testing"

	"github.com/jmervine/hipchat-sns-relay/config"

	. "github.com/jmervine/hipchat-sns-relay/Godeps/_workspace/src/gopkg.in/jmervine/GoT.v1"
)

var hc *Hipchat
var cfg = config.Parse([]string{"app", "-t", "token", "-r", "room"})

func TestNew(T *testing.T) {
	hc, err := New(cfg)

	Go(T).AssertNil(err)
	Go(T).RefuteNil(hc.client)
	Go(T).RefuteNil(hc.request)

	Go(T).AssertEqual(hc.room, "room")
}
