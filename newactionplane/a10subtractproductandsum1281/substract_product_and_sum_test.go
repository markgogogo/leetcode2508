package a10subtractproductandsum1281

import "testing"

func TestSubtractProductAndSum(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Example1", 234, 15},  // Product: 2*3*4=24, Sum: 2+3+4=9, Result: 24-9=15
		{"Example2", 4421, 21}, // Product: 4*4*2*1=32, Sum: 4+4+2+1=11, Result: 32-11=21
		{"SingleDigit", 5, 0},  // Product: 5, Sum: 5, Result: 5-5=0
		{"ZeroInput", 0, 0},    // Edge case: Product and Sum are 0
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subtractProductAndSum(tt.input)
			if result != tt.expected {
				t.Errorf("subtractProductAndSum(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
