package lib

import (
	"fmt"
	"math"
)

func GetGameClock(gameTime float64, startTime float64) string {
	difference := gameTime - startTime
	minuteDec := difference / 60
	minutes := math.Floor(minuteDec)
	seconds := round((minuteDec - minutes) * 60)

	minutesStr := fmt.Sprintf("%.2d", int(minutes))
	secondsStr := fmt.Sprintf("%.2d", int(seconds))

	return minutesStr + ":" + secondsStr
}

func round(a float64) float64 {
	if a < 0 {
		return math.Ceil(a - 0.5)
	}
	return math.Floor(a + 0.5)
}
