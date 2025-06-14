package utils

import (
	"regexp"
)

var (
	phoneRegExp = regexp.MustCompile(`^\+\d{10,15}$`)
	emailRegExp = regexp.MustCompile(`^(?:[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]{2,}(?:\.[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])$`)
	nameRegExp  = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9\s\-',.]{0,98}[A-Za-z0-9.]$`)
)

// Checks if the input string matches the expected phone number format.
// Used to ensure phone numbers in requests are valid before processing.
func ValidatePhoneNumber(value string) bool {
	return phoneRegExp.MatchString(value)
}

// Checks if the input string matches the expected email address format.
// Used to ensure emails in requests are valid before processing.
func ValidateEmail(value string) bool {
	return emailRegExp.MatchString(value)
}

// Checks if the input name is valid and normalizes it for consistency.
// Used to ensure names in requests are valid and human friendly before further processing.
func ValidateAndNormalizeName(name *string) bool {
	if name == nil {
		return false
	}

	normalized := regexp.MustCompile(`\s+`).ReplaceAllString(*name, " ")
	normalized = regexp.MustCompile(`^\s+|\s+$`).ReplaceAllString(normalized, "")

	if !nameRegExp.MatchString(normalized) {
		return false
	}

	if regexp.MustCompile(`[,.]{2,}`).MatchString(normalized) {
		return false
	}

	if len(normalized) <= 2 || len(normalized) > 32 {
		return false
	}

	*name = normalized

	return true
}
