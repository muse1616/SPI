package main

//brute force
// func minSubArrayLen(target int, nums []int) int {
//     ans:=math.MaxInt32
//     for i:=0;i<len(nums);i++{
//         sum:=0
//         for j:=i;j<len(nums);j++{
//             sum+=nums[j]
//             if sum>=target{
//                 ans = min(ans,j-i+1)
//                 break
//             }
//         }
//     }
//     if ans == math.MaxInt32{
//         return 0
//     }
//     return ans
// }
// func min(a,b int)int{
//     if a<b{
//         return a
//     }
//     return b
// }

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt32
	//双指针 left,right
	left, right := 0, 0
	// right先右移直到大于target 然后left右移直到小于left 更新答案 循环
	sum := 0
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			ans = min(ans, right-left+1)
			// left左移 减去left
			sum -= nums[left]
			left++
		}
		right++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
