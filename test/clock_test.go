package test

import (
	"dota2parser/lib"
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
