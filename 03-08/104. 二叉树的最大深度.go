package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	ans := 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, k int) {
		if root == nil {
			ans = max(ans, k)
			return
		}
		k++
		dfs(root.Left, k)
		dfs(root.Right, k)
	}
	dfs(root, 0)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
