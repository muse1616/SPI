package main

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	// 双指针
	left, right := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[left]*nums[left] >= nums[right]*nums[right] {
			ans[i] = nums[left] * nums[left]
			left++
		} else {
			ans[i] = nums[right] * nums[right]
			right--
		}
	}
	return ans
}
