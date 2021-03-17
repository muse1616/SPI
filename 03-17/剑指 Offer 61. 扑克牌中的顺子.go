package main

// 1.除了大小王 不能重复
// 2.max-min<5
func isStraight(nums []int) bool {
	hash_set := make(map[int]bool, 0)
	var (
		Max = -1
		Min = 20
	)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if _, ok := hash_set[nums[i]]; ok {
			return false
		}
		hash_set[nums[i]] = true
		Max = max(Max, nums[i])
		Min = min(Min, nums[i])
	}
	return Max-Min < 5
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
