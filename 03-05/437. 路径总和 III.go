package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 求出所有路径 然后判断路径中的序列是否包含连续的一部分 和为sum
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var pathFrom func(*TreeNode, int) int
	pathFrom = func(root *TreeNode, sum int) (cnt int) {
		if root == nil {
			return
		}
		if root.Val == sum {
			cnt++
		}
		return cnt + pathFrom(root.Left, sum-root.Val) + pathFrom(root.Right, sum-root.Val)
	}
	// 每个节点都可以当成第一个节点 所以左右子树递归外层函数
	return pathFrom(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
