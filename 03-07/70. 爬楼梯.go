package main

// 递归超时
// func climbStairs(n int) int {
//     if n <= 0{
//         return 0
//     }
//     if n == 1{
//         return 1
//     }
//     if n== 2{
//         return 2
//     }
//     return climbStairs(n-1)+climbStairs(n-2)
// }

// 动态规划
// func climbStairs(n int)int{
//     if n <=2 {
//         return n
//     }
//     dp:=make([]int,n+1)
//     dp[1] = 1
//     dp[2] = 2
//     for i:=3;i<=n;i++{
//         dp[i] = dp[i-1]+dp[i-2]
//     }
//     return dp[n]
// }
// 简化
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp_1 := 1
	dp_2 := 2
	dp_3 := 3
	for i := 3; i <= n; i++ {
		dp_3 = dp_1 + dp_2
		dp_1 = dp_2
		dp_2 = dp_3
	}
	return dp_3
}
