package lib

import (
	"fmt"
	"math"
	"strings"
)

func GetGameClock(gameTime float64, startTime float64) string {
	difference := gameTime - startTime
	minuteDec := difference / 60
	var minutes float64

	if minuteDec > 0 {
		minutes = math.Floor(minuteDec)
	} else {
		minutes = math.Ceil(minuteDec)
	}
	seconds := round((minuteDec - minutes) * 60)

	minutesStr := fmt.Sprintf("%.2d", int(minutes))
	secondsStr := fmt.Sprintf("%.2d", int(seconds))

	appended := minutesStr + ":" + secondsStr

	if strings.Index(appended, "-") > -1 {
		return "-" + strings.Replace(appended, "-", "", -1)
	}

	return appended
}

func round(a float64) float64 {
	if a < 0 {
		return math.Ceil(a - 0.5)
	}
	return math.Floor(a + 0.5)
}
