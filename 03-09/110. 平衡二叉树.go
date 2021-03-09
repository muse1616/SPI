package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return root == nil || (abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right))
}

// 求深度
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 左右max+1
	return max(height(root.Left), height(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 返回绝对值
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}
