package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}
	var dfs func(*TreeNode, []int)
	dfs = func(root *TreeNode, current []int) {
		if root == nil {
			return
		}
		current = append(current, root.Val)
		// 叶节点 求和
		if root.Left == nil && root.Right == nil {
			ans += sum(current)
			return
		}
		dfs(root.Left, append([]int{}, current...))
		dfs(root.Right, append([]int{}, current...))
	}
	dfs(root, []int{})
	return ans
}
func sum(arr []int) (s int) {
	for i := 0; i < len(arr); i++ {
		base := 1
		for j := len(arr) - i - 1; j > 0; j-- {
			base *= 10
		}
		s += arr[i] * base
	}
	return s
}
