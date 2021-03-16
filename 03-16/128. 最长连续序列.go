package main

func longestConsecutive(nums []int) int {
	// Hash去重
	nums_set := make(map[int]bool, 0)
	for _, v := range nums {
		nums_set[v] = true
	}
	// 当前长度 最大长度
	currentLen, maxLen := 0, 0

	for k := range nums_set {
		// 如果存在前驱 直接跳过
		if _, ok := nums_set[k-1]; ok {
			continue
		}
		// 不存在前驱向后遍历
		currentLen = 1
		for nums_set[k+1] {
			currentLen++
			k++
		}
		maxLen = max(maxLen, currentLen)
	}
	return maxLen
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
