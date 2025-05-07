package utils

import (
	"strconv"
	"strings"
)

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

func SplitString(value, separator string) []string {
	return strings.Split(value, separator)
}

func StringToNumber[T Number](value string) T {
	if integer, err := strconv.Atoi(value); err == nil {
		return any(integer).(T)
	}

	if decimal, err := strconv.ParseFloat(value, 64); err == nil {
		return any(decimal).(T)
	}

	var zero T
	return zero
}
