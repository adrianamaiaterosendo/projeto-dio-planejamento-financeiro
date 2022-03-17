package utils

import (
	"time"
)

const layout = "2006-01-02T15:04:05"

//StringToTime - Converte string em data
func StringToTime(value string) (time.Time, error) {
	convertedTime, err := time.Parse(layout, value)
	if err != nil {
		return time.Time{}, err
	}
	return convertedTime, nil
}
