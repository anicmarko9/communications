package utils

import (
	"testing"
)

// The structure for RegExp validation test cases.
type regexpTestCase struct {
	name  string
	input string
	want  bool
}

// Checks if the input string matches the expected phone number format.
func TestValidatePhoneNumber(t *testing.T) {
	tests := []regexpTestCase{
		{"Valid phone number", "+12345678901", true},
		{"Missing plus sign", "1234567890", false},
		{"Too short (less than 10 digits)", "+123", false},
		{"Too long (more than 15 digits)", "+1234567890123456", false},
		{"Invalid characters", "+1234abc8901", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePhoneNumber(tt.input); got != tt.want {
				t.Errorf("ValidatePhoneNumber(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

// Checks if the input string matches the expected email format.
func TestValidateEmail(t *testing.T) {
	tests := []regexpTestCase{
		{"Valid email", "test@example.com", true},
		{"Missing @ symbol", "test.example.com", false},
		{"Missing domain", "test@", false},
		{"Invalid characters", "test@exa!mple.com", false},
		{"Multiple @ symbols", "test@@example.com", false},
		{"Missing local part", "@example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateEmail(tt.input); got != tt.want {
				t.Errorf("ValidateEmail(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

// Convert a string to a pointer
func strPtr(s string) *string {
	return &s
}

// Checks if the input name is valid and normalizes it for consistency.
func TestValidateAndNormalizeName(t *testing.T) {
	tests := []struct {
		name   string
		input  *string
		want   bool
		output string
	}{
		{"Valid name", strPtr("John Doe"), true, "John Doe"},
		{"Name with extra spaces", strPtr("  John   Doe  "), true, "John Doe"},
		{"Name with leading/trailing tabs", strPtr("\tJohn Doe\t"), true, "John Doe"},
		{"Name with mixed whitespace", strPtr("  John\tDoe\n"), true, "John Doe"},
		{"Name with allowed punctuation", strPtr("John, Doe"), true, "John, Doe"},
		{"Name with dash", strPtr("John-Doe"), true, "John-Doe"},
		{"Name with apostrophe", strPtr("O'Connor"), true, "O'Connor"},
		{"Name with single period", strPtr("John. Doe"), true, "John. Doe"},
		{"Name with consecutive spaces", strPtr("John     Doe"), true, "John Doe"},
		{"Name with numbers", strPtr("John123"), true, "John123"},
		{"Name with period at end", strPtr("John."), true, "John."},
		{"Name with multiple spaces and punctuation", strPtr("  John   ,   Doe  "), true, "John , Doe"},
		{"Nil name", nil, false, ""},
		{"Name with extra spaces and symbols", strPtr("   _   John  _  Doe   _   "), false, ""},
		{"Name with invalid characters", strPtr("John@Doe"), false, ""},
		{"Name too short", strPtr("J"), false, ""},
		{"Name too long", strPtr("A very long name that exceeds the maximum length of ninety-eight characters, which is not allowed."), false, ""},
		{"Name with multiple consecutive punctuation", strPtr("!John.. ,Doe -_"), false, ""},
		{"Name with only spaces", strPtr("     "), false, ""},
		{"Name with only punctuation", strPtr("..."), false, ""},
		{"Name with underscore at start", strPtr("_John"), false, ""},
		{"Name with underscore at end", strPtr("John_"), false, ""},
		{"Name with dash at end", strPtr("John-"), false, ""},
		{"Name with comma at end", strPtr("John,"), false, ""},
		{"Name with period at start", strPtr(".John"), false, ""},
		{"Name with dash at start", strPtr("-John"), false, ""},
		{"Name with comma at start", strPtr(",John"), false, ""},
		{"Name with apostrophe at start", strPtr("'John"), false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var inputCopy *string
			if tt.input != nil {
				val := *tt.input
				inputCopy = &val
			}
			got := ValidateAndNormalizeName(inputCopy)

			if got != tt.want {
				t.Errorf("ValidateAndNormalizeName(%v) = %v, want %v", tt.input, got, tt.want)
			}

			if got && inputCopy != nil {
				expected := tt.output
				actual := *inputCopy
				if actual != expected {
					t.Errorf("ValidateAndNormalizeName(%v) normalized to %q, want %q", tt.input, actual, expected)
				}
			}
		})
	}
}
