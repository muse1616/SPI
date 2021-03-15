package main

// dp[i][j] 表示 A[i:] 和 B[j:] 的最长公共前缀，那么答案即为所有 dp[i][j] 中的最大值。如果 A[i] == B[j]，那么 dp[i][j] = dp[i + 1][j + 1] + 1，否则 dp[i][j] = 0。
func findLength(A []int, B []int) int {
	dp := make([][]int, len(A)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(B)+1)
	}
	maxLength := 0
	for i := len(A) - 1; i >= 0; i-- {
		for j := len(B) - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			maxLength = max(maxLength, dp[i][j])
		}
	}

	return maxLength
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
