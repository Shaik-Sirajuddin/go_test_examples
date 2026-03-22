package math

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{-1, -1, -2},
		{100, 200, 300},
	}

	for _, tt := range tests {
		got := Add(tt.a, tt.b)
		if got != tt.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{10, 5, 5},
		{0, 0, 0},
		{5, 10, -5},
		{-1, -1, 0},
		{100, 50, 50},
	}

	for _, tt := range tests {
		got := Subtract(tt.a, tt.b)
		if got != tt.expected {
			t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
		}
	}
}
