package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	var dfs func(root *TreeNode, currentSum int, path []int)
	dfs = func(root *TreeNode, currentSum int, path []int) {
		if root == nil {
			return
		}
		currentSum += root.Val
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil && currentSum == targetSum {
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs(root.Left, currentSum, append([]int{}, path...))
		dfs(root.Right, currentSum, append([]int{}, path...))
	}
	dfs(root, 0, []int{})
	return
	// path := []int{}
	// var dfs func(*TreeNode, int)
	// dfs = func(root *TreeNode, restSum int) {
	//     if root == nil {
	//         return
	//     }
	//     restSum -= root.Val
	//     path = append(path, root.Val)
	//     defer func() { path = path[:len(path)-1] }()
	//     if root.Left == nil && root.Right == nil && restSum == 0 {
	//         ans = append(ans, append([]int(nil), path...))
	//         return
	//     }
	//     dfs(root.Left, restSum)
	//     dfs(root.Right, restSum)
	// }
	// dfs(root, targetSum)
	// return

}
