package a11poweroftwo231

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Power of two: 1", 1, true},
		{"Power of two: 2", 2, true},
		{"Power of two: 4", 4, true},
		{"Not a power of two: 3", 3, false},
		{"Not a power of two: 0", 0, false},
		{"Not a power of two: -2", -2, false},
		{"Power of two: 8", 8, true},
		{"Not a power of two: 6", 6, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPowerOfTwo(tt.input)
			if result != tt.expected {
				t.Errorf("isPowerOfTwo(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
