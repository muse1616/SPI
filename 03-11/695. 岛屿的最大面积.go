package main

func maxAreaOfIsland(grid [][]int) (ans int) {
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		// 边界判断
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
			return 0
		}
		if grid[i][j] == 0 {
			return 0
		} else {
			grid[i][j] = 0
			return 1 + dfs(i-1, j) + dfs(i+1, j) + dfs(i, j+1) + dfs(i, j-1)
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
