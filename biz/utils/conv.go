package utils

import (
	"strconv"
	"time"
)

func ParseString2Uint64(s string) (uint64, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func ParseUint642String(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func ParseTime2DateString(t time.Time) string {
	return t.Format("2006-01-02")
}
