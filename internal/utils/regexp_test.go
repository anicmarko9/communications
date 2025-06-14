package utils

import "testing"

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
