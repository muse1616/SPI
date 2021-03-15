package main

func searchRange(nums []int, target int) []int {
	start, end := -1, -1
	// 先找起始点
	left, right := 0, len(nums)-1
	for left <= right {
		mid := right - (right-left+1)/2
		if nums[mid] == target {
			start = mid
			// 左边继续
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	left, right = 0, len(nums)-1
	for left <= right {
		mid := right - (right-left+1)/2
		if nums[mid] == target {
			end = mid
			// 右边继续
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return []int{start, end}
}
