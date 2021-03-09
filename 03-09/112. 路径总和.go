package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	// if root == nil{
	//     return false
	// }
	// ans:=false
	// var dfs func(root *TreeNode,currentSum int)
	// dfs = func(root *TreeNode,currentSum int){
	//     if root == nil{
	//         return
	//     }
	//     currentSum +=root.Val
	//     if root.Left == nil && root.Right == nil && currentSum == targetSum{
	//         ans = true
	//         return
	//     }
	//     dfs(root.Left,currentSum)
	//     dfs(root.Right,currentSum)
	// }
	// dfs(root,0)
	// return ans

	if root == nil{
		return false
	}
	if root.Left == nil && root.Right == nil && targetSum == root.Val{
		return true
	}
	return hasPathSum(root.Left,targetSum-root.Val)||hasPathSum(root.Right,targetSum-root.Val)
}