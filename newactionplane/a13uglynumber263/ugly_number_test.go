package a13uglynumber263

import "testing"

func TestIsUgly(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Negative number", -6, false},
		{"Zero", 0, false},
		{"One", 1, true},
		{"Ugly number 6", 6, true},
		{"Ugly number 8", 8, true},
		{"Ugly number 30", 30, true},
		{"Non-ugly number 14", 14, false},
		{"Non-ugly number 7", 7, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isUgly2(tt.input)
			if result != tt.expected {
				t.Errorf("isUgly(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
