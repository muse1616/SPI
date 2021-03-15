package main

// 双指针
func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for nums[left]&1 == 1 && left < right {
			left++
		}
		for nums[right]&1 == 0 && left < right {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}
