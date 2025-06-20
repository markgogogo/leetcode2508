package a05xoroperationinanarray1486

import "testing"

func TestXorOperation(t *testing.T) {
	tests := []struct {
		n        int
		start    int
		expected int
	}{
		{5, 0, 8},
		{4, 3, 8},
		{1, 7, 7},
		{10, 5, 2},
	}

	for _, test := range tests {
		result := xorOperation(test.n, test.start)
		if result != test.expected {
			t.Errorf("xorOperation(%d, %d) = %d; expected %d", test.n, test.start, result, test.expected)
		}
	}
}
