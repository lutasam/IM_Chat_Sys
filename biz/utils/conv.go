package utils

import "strconv"

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
