package main

func permute(nums []int) (ans [][]int) {
	var backtrace func([]int, int)
	// 回溯 原数组 当前序号
	backtrace = func(nums []int, index int) {
		if index == len(nums) {
			ans = append(ans, append([]int{}, nums...))
			return
		}
		// 交换
		for i := index; i < len(nums); i++ {
			nums[i], nums[index] = nums[index], nums[i]
			backtrace(nums, index+1)
			// 换回来
			nums[i], nums[index] = nums[index], nums[i]
		}

	}
	backtrace(nums, 0)
	return
}
