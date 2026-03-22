package stringutils

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},
		{"a", "a"},
		{"12345", "54321"},
		{"racecar", "racecar"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := Reverse(tt.input)
			if got != tt.expected {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{" ", false},
		{"a", false},
		{"not empty", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsEmpty(tt.input)
			if got != tt.expected {
				t.Errorf("IsEmpty(%q) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}
