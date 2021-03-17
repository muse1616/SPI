package main

func numWays(n int) int {
	dp_0, dp_1 := 1, 1
	for i := 2; i <= n; i++ {
		dp_0, dp_1 = dp_1, (dp_0+dp_1)%1000000007
	}
	return dp_1
}
