package a09adddigits258

import "testing"

func TestAddDigits(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Single digit", 5, 5},
		{"Two digits", 38, 2},
		{"Multiple digits", 123, 6},
		{"Boundary Value", 10, 1},
		{"Zero input", 0, 0},
		{"Large number", 987654321, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addDigits(tt.input)
			if result != tt.expected {
				t.Errorf("addDigits(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
