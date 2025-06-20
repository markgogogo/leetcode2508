package a15transposematrix

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			name: "Square matrix",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
		{
			name: "Rectangular matrix (more rows)",
			matrix: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			expected: [][]int{
				{1, 3, 5},
				{2, 4, 6},
			},
		},
		{
			name: "Rectangular matrix (more columns)",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expected: [][]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
		{
			name: "Single row matrix",
			matrix: [][]int{
				{1, 2, 3},
			},
			expected: [][]int{
				{1},
				{2},
				{3},
			},
		},
		{
			name: "Single column matrix",
			matrix: [][]int{
				{1},
				{2},
				{3},
			},
			expected: [][]int{
				{1, 2, 3},
			},
		},
		{
			name:     "empty matrix",
			matrix:   [][]int{},
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transpose(tt.matrix)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("transpose() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
