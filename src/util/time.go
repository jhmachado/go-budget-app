package util

import "time"

var layout = "2006-01-02T15:04:05"

func ConvertStringToTime(date string) time.Time {
	formattedDate, _ := time.Parse(layout, date)
	return formattedDate
}
