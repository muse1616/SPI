package main

// 空间复杂度O(n)
// func firstMissingPositive(nums []int) int {
//     hashTable:=make(map[int]bool,0)
// 	for i:=0;i<len(nums);i++{
// 		hashTable[nums[i]] = true
// 	}
// 	for i:=1;;i++{
// 		if _,ok:=hashTable[i];!ok{
// 			return i
// 		}
// 	}
// }
// 常数级别空间
// 对于长度是N的nums 答案只有可能是[1,N+1]
func firstMissingPositive(nums []int) int {
	// 小于等于0的数字全部置为N+1
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 1
		}
	}
	for i := 0; i < len(nums); i++ {
		// 此时这个数字可能已经被置为负数 先还原
		now := nums[i]
		if nums[i] < 0 {
			now = -nums[i]
		}
		// 根据当前x 将index为[x-1]的数字变为负数 如果本来就是负数 就不用管
		if now < len(nums)+1 {
			if nums[now-1] > 0 {
				// 负数标记
				nums[now-1] = -nums[now-1]
			}
		}
	}
	// 寻找第一个负数的标志
	ans := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			ans = i + 1
			break
		}
	}
	return ans
}
