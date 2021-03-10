package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// 保证第一个数组长度较小
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	// 分割线左边的元素
	// 使用除法的取整性质 可以避免考虑总长度为奇数还是偶数的两种情况
	totalLeft := (m + n + 1) / 2
	left, right := 0, m
	// 需要在nums1确定划分线
	// 使得 nums1[i-1]<=nums2[j] && nums2[j-1] <= nums1[i]
	for left < right {
		i := left + (right-left+1)/2
		j := totalLeft - i
		if nums1[i-1] > nums2[j] {
			// 说明划分线太靠右边了
			right = i - 1
		} else {
			// 说明划分线太靠左边了
			left = i
		}
	}
	// 此时需要判断划分的几种极端情况
	// 此时考虑四个数字
	var (
		nums1LeftMax  int
		nums1RightMin int
		nums2LeftMax  int
		nums2RightMin int
	)
	i, j := left, totalLeft-left
	// 第一种情况 nums1的划分线划在最左边 i = 0
	// 比如 | 4 5 6
	//      - - -- --
	//        1 2 3 |
	if i == 0 {
		// 使其在比较最大值的时候默认被去掉
		nums1LeftMax = math.MinInt32
	} else {
		// 左边这个值
		nums1LeftMax = nums1[i-1]
	}
	// 上面的划分在最右边
	if i == m {
		nums1RightMin = math.MaxInt32
	} else {
		nums1RightMin = nums1[i]
	}
	// 第二个数组同理
	if j == 0 {
		nums2LeftMax = math.MinInt32
	} else {
		nums2LeftMax = nums2[j-1]
	}
	if j == n {
		nums2RightMin = math.MaxInt32
	} else {
		nums2RightMin = nums2[j]
	}

	// 总长度为奇数
	if (m+n)%2 == 1 {
		// 左边的最大值
		return float64(max(nums1LeftMax, nums2LeftMax))
	}
	// fmt.Println(max(nums1LeftMax,nums2LeftMax)+min(nums1RightMin,nums2RightMin))
	// fmt.Println((float64)(max(nums1LeftMax,nums2LeftMax)+min(nums1RightMin,nums2RightMin))/2.0)
	// // 总长度为偶数 左边的最大值 和 右边的最小值的平均值
	return float64((float64)(max(nums1LeftMax, nums2LeftMax)+min(nums1RightMin, nums2RightMin)) / 2.0)
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
