package hipchat

import (
	"testing"

	"github.com/jmervine/hipchat-sns-relay/config"

	. "github.com/jmervine/hipchat-sns-relay/Godeps/_workspace/src/gopkg.in/jmervine/GoT.v1"
)

var hc *Hipchat
var cfg = config.Parse([]string{"app", "-t", "token", "-r", "room"})

func TestNew(T *testing.T) {
	hc = New(cfg)

	Go(T).AssertEqual(hc.params.Get("room_id"), "room")
	Go(T).AssertEqual(hc.auth.Get("auth_token"), "token")
	Go(T).AssertEqual(hc.endpoint, "https://api.hipchat.com/v1/rooms/message")
}
