package main

// 归并排序
func reversePairs(nums []int) int {
	var mergeSort func(int, int) int
	mergeSort = func(start int, end int) int {
		if start >= end {
			return 0
		}
		middle := start + (end-start)/2
		cnt := mergeSort(start, middle) + mergeSort(middle+1, end)
		// 临时数组开辟
		tmp := []int{}
		i, j := start, middle+1
		for i <= middle && j <= end {
			if nums[i] <= nums[j] {
				// 加入左边这个数字的时候 产生逆序对 即比之前加入的右边的数字大
				tmp = append(tmp, nums[i])
				cnt += j - (middle + 1)
				i++
			} else {
				tmp = append(tmp, nums[j])
				j++
			}
		}
		// 左边的数组贡献逆序对
		for ; i <= middle; i++ {
			tmp = append(tmp, nums[i])
			cnt += end - middle
		}
		for ; j <= end; j++ {
			tmp = append(tmp, nums[j])
		}
		// 临时数组还原
		for i := start; i <= end; i++ {
			nums[i] = tmp[i-start]
		}
		return cnt
	}
	return mergeSort(0, len(nums)-1)
}
