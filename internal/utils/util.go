package utils

import (
	"strconv"
	"strings"
	"time"
)

// Generic interface for all common numeric types.
// Used to allow type-safe conversions and operations in generic functions.
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Common size constants for working with bytes.
const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

// Splits a string into a slice using the provided separator.
// Useful for parsing delimited input values.
func SplitString(value, separator string) []string {
	return strings.Split(value, separator)
}

// Converts a string to a specified numeric type (int, float, etc.).
// Useful for converting database string results to Go numeric types in a type-safe way.
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

// Returns the current UTC time in RFC3339 format.
// Used for consistent timestamping in API responses and logs.
func GetCurrentTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}
