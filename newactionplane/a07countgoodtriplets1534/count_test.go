package a07countgoodtriplets1534

import "testing"

func TestCountGoodTriplets(t *testing.T) {
	tests := []struct {
		arr      []int
		a, b, c  int
		expected int
	}{
		{[]int{3, 0, 1, 1, 9, 7}, 7, 2, 3, 4},
		{[]int{1, 1, 2, 2, 3}, 0, 0, 1, 0},
		{[]int{1, 2, 3, 4}, 1, 1, 1, 0},
		{[]int{1, 1, 1}, 0, 0, 0, 1},
		{[]int{}, 1, 1, 1, 0},
	}

	for _, test := range tests {
		result := countGoodTriplets(test.arr, test.a, test.b, test.c)
		if result != test.expected {
			t.Errorf("For arr=%v, a=%d, b=%d, c=%d; expected %d, got %d",
				test.arr, test.a, test.b, test.c, test.expected, result)
		}
	}
}
