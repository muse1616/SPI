package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 辅助数组
// func flatten(root *TreeNode)  {
//     var arr []*TreeNode
//     var preorder func(*TreeNode)
//     preorder = func(root *TreeNode){
//         if root == nil{
//             return
//         }
//         arr = append(arr,root)
//         preorder(root.Left)
//         preorder(root.Right)
//     }
//     preorder(root)
//     for i:=0;i<len(arr);i++{
//         arr[i].Left =nil
//         if i == 0{
//             continue
//         }
//         arr[i-1].Right = arr[i]
//     }
// }

func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}
