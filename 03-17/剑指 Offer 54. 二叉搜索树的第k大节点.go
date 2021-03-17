package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var skip int
var res int

func kthLargest(root *TreeNode, k int) int {
	skip = k
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Right)
		skip--
		if skip == 0 {
			res = root.Val
			return
		}
		dfs(root.Left)
	}
}
