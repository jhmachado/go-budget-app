package util

import "testing"

func TestConvertStringToTime(t *testing.T) {
	var convertedTime = ConvertStringToTime("2019-02-12T10:10:10")

	if convertedTime.Year() != 2019 {
		t.Errorf("Get the wrong year from converted time")
	}

	if convertedTime.Month() != 02 {
		t.Errorf("Get the wrong month from converted time")
	}

	if convertedTime.Day() != 12 {
		t.Errorf("Get the wrong day from converted time")
	}
}
