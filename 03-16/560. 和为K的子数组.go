package main

// 前缀和+hash优化
func subarraySum(nums []int, k int) (cnt int) {
	var (
		pre int
		m   = make(map[int]int, 0)
	)
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		// 前缀和-k存在
		if _, ok := m[pre-k]; ok {
			cnt += m[pre-k]
		}
		m[pre] += 1
	}
	return
}
