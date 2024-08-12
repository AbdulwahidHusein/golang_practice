package validation

import (
	"testing"
)

func TestEmailValidator(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"", false},
		{"abc", false},
		{"abc@", false},
		{"abc@abc", false},
		{"abc@abc.", false},
		{"abc@abccom", false},
		{"abc@abccom.", false},
		{"abc@abc.com", true},
		{"a@a.a", false},
		{"a@a.a.bc", true},
		{"a@a.a.com", true},
	}

	for _, sample := range tests {
		t.Run(sample.email, func(t *testing.T) {
			result := IsValidEmail(sample.email)
			if result != sample.valid {
				t.Errorf("IsValidEmail(%q) = %v; want %v", sample.email, result, sample.valid)
			}
		})
	}

}

func TestPasswordValidator(t *testing.T) {
	tests := []struct {
		password string
		valid    bool
	}{
		{"", false},
		{"nsdjkfnsdjk", false},
		{"Abcsdn34", false},
		{"abc12abc", false},
		{"abc@abcA", false},
		{"abc@ab9ccom", false},
		{"abc@abAcco@m.", false},
		{"abc@abcWe3.com", true},
		{"1A@Bb", false},
		{"a@a.a.bE#0c", true},
		{"a@a.a.com", false},
	}

	for _, sample := range tests {
		t.Run(sample.password, func(t *testing.T) {
			result := IsValidPassword(sample.password)
			if result != sample.valid {
				t.Errorf("IsValidPassword(%q) = %v; want %v", sample.password, result, sample.valid)
			}
		})
	}

}
