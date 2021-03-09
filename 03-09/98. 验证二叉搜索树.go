package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归即可
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l, r := true, true
	if root.Left != nil {
		l = root.Val > maxFunc(root.Left)
	}
	if root.Right != nil {
		r = root.Val < minFunc(root.Right)
	}
	return l && r && isValidBST(root.Left) && isValidBST(root.Right)
}

// 返回其中最大的节点
func maxFunc(root *TreeNode) (ans int) {
	ans = math.MinInt32
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = max(ans, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return
}
func minFunc(root *TreeNode) (ans int) {
	ans = math.MaxInt32
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = min(ans, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
