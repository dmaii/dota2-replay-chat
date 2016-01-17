package test

import (
	"github.com/suisha/dota2-replay-chat/lib"
	"testing"
)

func TestGetGameClock(t *testing.T) {
	var a float64 = 450
	var b float64 = 200

	str := lib.GetGameClock(a, b)

	if str != "04:10" {
		t.Fail()
	}
}
