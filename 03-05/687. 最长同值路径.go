package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func longestUnivaluePath(root *TreeNode) (ans int) {
	// 从一个根向两边递归
	var dfs func(*TreeNode)int
	dfs = func(root *TreeNode) int {
		if root==nil{
			return 0
		}
		// 左子树
		leftSub := dfs(root.Left)
		// 右子树
		rightSub := dfs(root.Right)
		// 当前路径长度
		l,r:=0,0
		// 如果左子树不为空 且左节点等于当前节点
		if root.Left!=nil && root.Left.Val == root.Val {
			l += (leftSub+1)
		}
		// 同上
		if root.Right!=nil && root.Right.Val == root.Val{
			r += (rightSub +1)
		}
		ans = max(ans,l+r)
		return max(l,r)
	}

	dfs(root)
	return
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}