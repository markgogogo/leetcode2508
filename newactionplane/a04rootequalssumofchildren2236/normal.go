package a04rootequalssumofchildren2236

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool { // TODO: 扩展到n个节点？
	return root.Val == root.Left.Val+root.Right.Val
}
