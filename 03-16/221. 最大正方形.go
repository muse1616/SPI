package main

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}
	// 边长
	maxBound := 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			// bound
			if i == 0 || j == 0 {
				if matrix[i][j] == '1' {
					dp[i][j] = 1
					maxBound = max(maxBound, dp[i][j])
				}
				continue
			}
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				maxBound = max(maxBound, dp[i][j])
			}
		}
	}
	return maxBound * maxBound
}
func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= c && b <= a {
		return b
	}
	return c
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
