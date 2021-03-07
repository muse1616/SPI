package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	var dfs func(*TreeNode,*TreeNode)bool
	dfs = func(left *TreeNode,right *TreeNode)bool{
		if left == nil && right == nil{
			return true
		}
		// 因为已经判断过两个都是nil返回true 这里只有一个为nil
		if left == nil || right == nil{
			return false
		}
		return left.Val == right.Val && dfs(left.Left,right.Right) && dfs(left.Right,right.Left)
	}
	return dfs(root,root)
}