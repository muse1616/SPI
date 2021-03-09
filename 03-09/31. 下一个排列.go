package main

// 主要思路查找一个比当前序列大一个排列的序列
// 需要我们需要将一个左边的「较小数」与一个右边的「较大数」交换，以能够让当前排列变大，从而得到下一个排列。同时我们要让这个「较小数」尽量靠右，而「较大数」尽可能小。当交换完成后，「较大数」右边的数需要按照升序重新排列。这样可以在保证新排列大于原来排列的情况下，使变大的幅度尽可能小。
func nextPermutation(nums []int) {
	// 从后向前搜索 a[i]<a[i+1]
	// 这时[i+1,n]一定是降序排列
	// 再从区间[i+1,n]中从后往前查找第一个大于a[i]的数字 这样这个数字就会在降序排列中最靠右
	// 也就满足了这个数字尽可能小
	// 然后[i+1,n]也会是降序 将这个序列利用双指针逆序变为升序即可

	// 标志变量 是否进行过操作
	flag := false
	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			flag = true
			// 在[i+1,n]中从后往前查找第一个大于nums[i]的数字
			j := len(nums) - 1
			for ; j >= i+1; j-- {
				if nums[j] > nums[i] {
					break
				}
			}
			// 交换 i j
			nums[i], nums[j] = nums[j], nums[i]
			// 逆序 [i+1,n]
			left, right := i+1, len(nums)-1
			for left < right {
				nums[left], nums[right] = nums[right], nums[left]
				left++
				right--
			}
			break
		}
	}

	// 未进行过操作 逆序
	if flag == false {
		for i := 0; i < len(nums)/2; i++ {
			nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
		}
	}
}
