package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// func inorderTraversal(root *TreeNode)(res []int) {
//     var helper func(root *TreeNode)
//     helper = func(root *TreeNode){
//         if root==nil{
//             return
//         }

//         helper(root.Left)
//          res = append(res,root.Val)
//         helper(root.Right)
//     }
//     helper(root)
//     return
// }
// 迭代
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
