package a01addtwointegers2235

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		num1     int
		num2     int
		expected int
	}{
		{12, 5, 17},
		{-10, 4, -6},
		{1, 2, 3},
		{-1, -1, -2},
		{0, 0, 0},
		{100, 200, 300},
		{-100, 100, 0},
		{9223372036854775807, 1, -9223372036854775808},
	}

	for _, test := range tests {
		result := sum(test.num1, test.num2)
		if result != test.expected {
			t.Errorf("sum(%d, %d) = %d; expected %d", test.num1, test.num2, result, test.expected)
		}
	}
}
