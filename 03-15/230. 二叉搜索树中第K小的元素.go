package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 中遍历二叉树?
func kthSmallest(root *TreeNode, k int) int {
	nodes := make([]int, 0)
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		nodes = append(nodes, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return nodes[k-1]
}
