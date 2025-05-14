package utils

import (
	"regexp"
)

func ValidatePhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+\d{10,15}$`)

	return re.MatchString(phone)
}
