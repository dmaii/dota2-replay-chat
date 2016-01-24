package lib_test

import (
	"github.com/suisha/dota2-replay-chat/lib"
	"testing"
)

func TestClockBase(t *testing.T) {
	var a float64 = 450
	var b float64 = 200

	str := lib.GetGameClock(a, b)

	if str != "04:10" {
		t.Fail()
	}
}

func TestClockNegative(t *testing.T) {
	var a float64 = 100
	var b float64 = 150

	str := lib.GetGameClock(a, b)

	if str != "-00:50" {
		t.Fail()
	}
}
