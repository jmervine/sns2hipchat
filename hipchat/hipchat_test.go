package hipchat

import (
	"testing"

	"github.com/jmervine/sns2hipchat/config"

	. "github.com/jmervine/sns2hipchat/Godeps/_workspace/src/gopkg.in/jmervine/GoT.v1"
)

var cfg = config.Parse([]string{"app", "-t", "token", "-r", "room"})

func TestNewV2(T *testing.T) {
	hc, err := NewV2(cfg)

	Go(T).AssertNil(err)
	Go(T).RefuteNil(hc)
	Go(T).RefuteNil(v2)
	Go(T).RefuteNil(hc.client)
	Go(T).RefuteNil(hc.request)

	Go(T).AssertEqual(hc, v2, "should set HipchatV2 in global namespace for reuse")
}

func TestV2_Post(t *testing.T) {
	t.Skip("TODO: stub v2.client.Room.Notification")

	// perhaps by stubbing the method to return http.Response in a
	// sensical and testable manner
}

func TestNewV1(T *testing.T) {
	hc, err := NewV1(cfg)

	Go(T).AssertNil(err)
	Go(T).RefuteNil(hc)
	Go(T).RefuteNil(v1)
	Go(T).RefuteNil(hc.client)
	Go(T).RefuteNil(hc.request)
	Go(T).RefuteNil(hc.auth)

	Go(T).AssertEqual(hc, v1, "should set HipchatV1 in global namespace for reuse")
}

func TestV1_Post(t *testing.T) {
	t.Skip("TODO: stub http.PostFrom")

	// perhaps by assigning httpPostFrom := http.PostForm in
	// the global context
}
