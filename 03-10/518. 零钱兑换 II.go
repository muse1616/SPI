package main

// 回溯超时
// func change(amount int, coins []int) (ans int) {
//     if amount == 0{
//         return 1
//     }
//     var backtrace func(int,int)
//     backtrace = func(amount int,index int){
//         if index == len(coins) || amount <= 0{
//             if amount == 0{
//                 ans++
//             }
//             return
//         }
//         backtrace(amount-coins[index],index)
//         backtrace(amount,index+1)
//     }
//     backtrace(amount,0)
//     return ans
// }

//dp
func change(amount int, coins []int) (ans int) {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, v := range coins {
		for x := v; x <= amount; x++ {
			dp[x] += dp[x-v]
		}
	}
	return dp[amount]
}
