package main

// O(n^2)dp
func lengthOfLIS(nums []int) (ans int) {
	if len(nums) == 1{
		return 1
	}
	dp:=make([]int,len(nums))
	dp[0] = 1
	ans = 1
	for i:=1;i<len(nums);i++{
		dp[i] = 1
		for j:=0;j<i;j++{
			if nums[j]<nums[i]{
				dp[i] = max(dp[i],dp[j]+1)
				ans = max(ans,dp[i])
			}
		}
	}
	return``
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
