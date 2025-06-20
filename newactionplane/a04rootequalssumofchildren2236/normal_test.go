package a04rootequalssumofchildren2236

import "testing"

func TestCheckTree(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected bool
	}{
		{&TreeNode{Val: 10, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}}, true},
		{&TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}, false},
		{&TreeNode{Val: 0, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 0}}, true},
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 1}}, true},
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}}, false},
	}

	for _, test := range tests {
		result := checkTree(test.root)
		if result != test.expected {
			t.Errorf("checkTree(%v) = %v; expected %v", test.root, result, test.expected)
		}
	}
}
