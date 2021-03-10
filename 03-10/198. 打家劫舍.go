package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		if i == 0 {
			dp[i] = nums[0]
		} else if i == 1 {
			dp[i] = max(nums[0], nums[1])
		} else {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		}

	}

	return dp[len(nums)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
