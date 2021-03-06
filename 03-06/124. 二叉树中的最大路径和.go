package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode)(ans int) {
	// 因为二叉树中有负数 所以默认的答案需要是个特别小的负数
	ans = -1*math.MaxInt32
	// 深度搜索辅助函数
	var dfs func(*TreeNode)int
	dfs = func(root *TreeNode)int{
		// 节点为空 返回 递归终点
		if root==nil{
			return 0
		}
		// 分别递归左右子树 只有当左右子树的最大路径是正数的时候 才会对以当前节点作为最高节点的路径产生最大值的贡献
		l:=max(dfs(root.Left),0)
		r:=max(dfs(root.Right),0)
		// 左子树最长路劲+当前节点+右子树最长路径 更新ans
		ans = max(l+r+root.Val,ans)
		// 返回当前值加左或者右子树最长的一条 因为路径只能是一条线
		return root.Val+max(l,r)
	}
	dfs(root)

	return ans
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
