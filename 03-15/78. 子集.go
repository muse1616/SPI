package main

func subsets(nums []int) (ans [][]int) {
	var backtrace func([]int, int)
	backtrace = func(current []int, index int) {
		if index == len(nums) {
			ans = append(ans, append([]int{}, current...))
			return
		}
		backtrace(append(current, nums[index]), index+1)
		backtrace(current, index+1)
	}
	backtrace([]int{}, 0)

	return
}
