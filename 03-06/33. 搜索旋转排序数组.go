package main

// 二分查找
func search(nums []int, target int) int {
	if len(nums) == 0{
		return -1
	}
	if len(nums) == 1{
		if target == nums[0] {
			return 0
		}
		return -1
	}
	// 左右
	l,r:=0,len(nums)-1
	for l<=r{
		// 防止溢出
		mid:=r-(r-l)/2
		if nums[mid] == target{
			return mid
		}
		// 左侧是否有序判断
		if nums[0]<=nums[mid]{
			if nums[0]<=target && nums[mid]>target{
				r = mid-1
			}else {
				l = mid+1
			}
		}else{
			// 右侧有序
			if nums[mid]<target && target<=nums[len(nums)-1]{
				l = mid+1
			}else {
				r = mid-1
			}
		}

	}


	return -1
}