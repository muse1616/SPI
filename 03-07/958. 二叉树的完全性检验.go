package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 层序遍历 如果读到空 后面还有节点 返回false
func isCompleteTree(root *TreeNode) bool {
	last:=false
	q := []*TreeNode{root}
	for len(q)>0{
		node:=q[0]
		q = q[1:]
		if node == nil {
			last = true
		}else{
			if last == true{
				return false
			}
			// 左右入列
			q = append(q,node.Left)
			q = append(q,node.Right)
		}
	}
	return true
}