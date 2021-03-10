package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row == 0 || col == 0 {
				dp[row][col] = 1
			} else {
				dp[row][col] = dp[row-1][col] + dp[row][col-1]
			}
		}
	}
	return dp[m-1][n-1]
}
