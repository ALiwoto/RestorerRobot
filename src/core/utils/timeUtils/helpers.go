package timeUtils

import (
	"strconv"
	"time"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/stringUtils"
)

// GenerateCurrentDateTime format of the date time will be dd/MM/yyyy HH:mm:ss
func GenerateCurrentDateTime() string {
	t := time.Now()

	str := stringUtils.MakeSureNum(t.Day(), 2) + "/"
	str += stringUtils.MakeSureNum(int(t.Month()), 2) + "/"
	str += stringUtils.MakeSureNum(t.Year(), 4) + " "
	str += stringUtils.MakeSureNum(t.Hour(), 2) + ":"
	str += stringUtils.MakeSureNum(t.Minute(), 2) + ":"
	str += stringUtils.MakeSureNum(t.Second(), 2)

	return str
}

// GenerateSuitableDateTime will format of the date time to dd-MM-yyyy HH-mm-ss
func GenerateSuitableDateTime() string {
	t := time.Now()

	str := stringUtils.MakeSureNum(t.Day(), 2) + "-"
	str += stringUtils.MakeSureNum(int(t.Month()), 2) + "-"
	str += stringUtils.MakeSureNum(t.Year(), 4) + "--"
	str += stringUtils.MakeSureNum(t.Hour(), 2) + "-"
	str += stringUtils.MakeSureNum(t.Minute(), 2) + "-"
	str += stringUtils.MakeSureNum(t.Second(), 2)

	return str
}

func GetPrettyTimeDuration(d time.Duration, shorten bool) string {
	var result string
	totalSeconds := int(d.Seconds())

	year := totalSeconds / (60 * 60 * 24 * 365)
	totalSeconds -= year * (60 * 60 * 24 * 365)

	month := totalSeconds / (60 * 60 * 24 * 30)
	totalSeconds -= month * (60 * 60 * 24 * 30)

	day := totalSeconds / (60 * 60 * 24)
	totalSeconds -= day * (60 * 60 * 24)

	hour := totalSeconds / (60 * 60)
	totalSeconds -= hour * (60 * 60)

	minute := totalSeconds / 60
	totalSeconds -= minute * 60

	seconds := totalSeconds

	yBool := year > 0
	mBool := month > 0 || yBool
	shorten = !mBool && shorten
	dBool := day > 0 || mBool
	hBool := hour > 0 || dBool
	if yBool {
		result += strconv.Itoa(year) + " year"
		if year > 1 {
			result += "s"
		}
		result += " "
	}
	if mBool {
		result += " " + strconv.Itoa(month) + " month"
		if month > 1 {
			result += "s"
		}
		result += " "
	}
	if dBool {
		result += strconv.Itoa(day)
		if shorten {
			result += "d"
		} else {
			result += " day"
			if day > 1 {
				result += "s"
			}
		}
		result += " "
	}
	if hBool {
		result += strconv.Itoa(hour)
		if shorten {
			result += "h"
		} else {
			result += " hour"
			if hour > 1 {
				result += "s"
			}
		}
		result += " "
	}
	result += strconv.Itoa(minute)
	if shorten {
		result += "m"
	} else {
		result += " minute"
		if minute > 1 {
			result += "s"
		}
	}

	result += " " + strconv.Itoa(seconds)
	if shorten {
		result += "s"
	} else {
		result += " second"
		if seconds > 1 {
			result += "s"
		}
	}
	return result
}
