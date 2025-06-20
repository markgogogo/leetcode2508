package a12powerofthree326

import "testing"

func TestIsPowerOfThree(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Power of three: 1", 1, true},
		{"Power of three: 3", 3, true},
		{"Power of three: 9", 9, true},
		{"Power of three: 27", 27, true},
		{"Not a power of three: 0", 0, false},
		{"Not a power of three: 2", 2, false},
		{"Not a power of three: 10", 10, false},
		{"Not a power of three: 28", 28, false},
		{"Not a power of three: -1", -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPowerOfThree1(tt.input)
			if result != tt.expected {
				t.Errorf("isPowerOfThree(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
