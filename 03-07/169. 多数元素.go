package main

// 不能排序
func majorityElement(nums []int) (candidate int) {
	cnt:=0
	for _,v:=range nums{
		if cnt == 0{
			candidate = v
		}
		if candidate == v{
			cnt++
		} else{
			cnt--
		}
	}
	return
}