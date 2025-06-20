package a03smallestevenmultiple2413

import "testing"

func TestSmallestEvenMultiple(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{5, 10},
		{6, 6},
		{1, 2},
		{2, 2},
		{15, 30},
	}

	for _, test := range tests {
		result := smallestEvenMultiple(test.input)
		if result != test.expected {
			t.Errorf("smallestEvenMultiple(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}
