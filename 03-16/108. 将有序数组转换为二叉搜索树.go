package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var buildTree func(int, int) *TreeNode
	buildTree = func(left, right int) *TreeNode {
		// 每次都选最中间的一个元素作为父节点
		if left > right {
			return nil
		}
		mid := right - (right-left)/2
		root := &TreeNode{Val: nums[mid]}
		root.Left = buildTree(left, mid-1)
		root.Right = buildTree(mid+1, right)
		return root
	}
	return buildTree(0, len(nums)-1)
}
