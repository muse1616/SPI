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
// func preorderTraversal(root *TreeNode) (ans []int) {
//     var preorder func(*TreeNode)
//     preorder = func(root *TreeNode){
//         if root==nil{
//             return
//         }
//         ans = append(ans,root.Val)
//         preorder(root.Left)
//         preorder(root.Right)
//     }
//     preorder(root)
//     return ans
// }
// 迭代
func preorderTraversal(root *TreeNode) (ans []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			ans = append(ans, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}
