package formatter

import (
	"testing"

	. "github.com/jmervine/sns2hipchat/Godeps/_workspace/src/gopkg.in/jmervine/GoT.v1"
)

func TestNew(T *testing.T) {
	f := New("basic")
	Go(T).RefuteNil(f)
}
