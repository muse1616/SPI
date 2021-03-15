package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 深度遍历 记录每一个节点的深度 放入hash
// 然后对于每一个节点，它对应这一层的可能宽度是 pos - left[depth] + 1
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(*TreeNode, int, int)
	ans := 1
	hashTable := make(map[int]int, 0)
	dfs = func(root *TreeNode, depth int, position int) {
		if root == nil {
			return
		}
		if v, ok := hashTable[depth]; ok {
			ans = max(ans, position-v+1)
		} else {
			hashTable[depth] = position
		}
		dfs(root.Left, depth+1, position*2)
		dfs(root.Right, depth+1, position*2+1)
	}
	dfs(root, 0, 0)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
