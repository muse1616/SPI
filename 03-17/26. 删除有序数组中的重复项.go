package main

// 传统双指针
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// 新数组中真正的下标
	i := 0
	for j := 1; j < len(nums); j++ {
		// 如果等于跳过
		if nums[j] == nums[i] {
			continue
		}
		// 不重复
		i++
		nums[i] = nums[j]
	}

	return i + 1
}

// 语言特性
// func removeDuplicates(nums []int) int {
//     if len(nums) == 0 || len(nums) == 1{
//         return  len(nums)
//     }
//     pre,cur:=0,1
//     for cur<len(nums){
//         if nums[cur] == nums[pre]{
//             // 删除nums[cur]
//             nums = append(nums[:cur],nums[cur+1:]...)
//         }else{
//             pre = cur
//             cur ++
//         }
//     }

//     return len(nums)
// }
