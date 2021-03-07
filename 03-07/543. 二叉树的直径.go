package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 左 右
		left := dfs(root.Left)
		right := dfs(root.Right)
		ans = max(ans, left+right+1)
		return 1 + max(left, right)
	}
	dfs(root)
	return ans - 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
