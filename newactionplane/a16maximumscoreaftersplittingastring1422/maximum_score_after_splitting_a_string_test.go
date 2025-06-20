package a16maximumscoreaftersplittingastring1422

import "testing"

func TestMaxScore(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"011101", 5},
		{"00111", 5},
		{"1111", 3},
		{"0000", 3},
		{"010101", 4},
		{"10", 0},
		{"01", 2},
	}

	for _, test := range tests {
		result := maxScore4(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d but got %d", test.input, test.expected, result)
		}
	}
}
