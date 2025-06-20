package a06numberofgoodpairs1512

import "testing"

func TestNumIdenticalPairs(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1, 1, 3}, 4},
		{[]int{1, 1, 1, 1}, 6},
		{[]int{1, 2, 3}, 0},
		{[]int{}, 0},
		{[]int{1}, 0},
	}

	for _, test := range tests {
		result := numIdenticalPairs3(test.nums)
		if result != test.expected {
			t.Errorf("numIdenticalPairs(%v) = %d; expected %d", test.nums, result, test.expected)
		}
	}
}
