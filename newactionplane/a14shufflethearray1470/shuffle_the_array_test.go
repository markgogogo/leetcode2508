package a14shufflethearray

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		n        int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{10, 1, 2, 3, 4, 5},
			n:        3,
			expected: []int{10, 3, 1, 4, 2, 5},
		},
		// {
		// 	name:     "Example 2",
		// 	nums:     []int{1, 2, 3, 4, 4, 3, 2, 1},
		// 	n:        4,
		// 	expected: []int{1, 4, 2, 3, 3, 2, 4, 1},
		// },
		// {
		// 	name:     "Example 3",
		// 	nums:     []int{1, 1, 2, 2},
		// 	n:        2,
		// 	expected: []int{1, 2, 1, 2},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := shuffle5_2(tt.nums, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("shuffle(%v, %d) = %v; want %v", tt.nums, tt.n, result, tt.expected)
			}
		})
	}
}
