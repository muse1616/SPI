package main

// O(n)
// func findPeakElement(nums []int) int {
//     if len(nums)==1{
//         return 0
//     }
//     for i:=0;i<len(nums);i++{
//         if i==0 {
//             if  nums[i]>nums[i+1]{
//                 return 0
//             }
//             continue
//         }
//         if i == len(nums)-1 {
//             if  nums[i]>nums[i-1]{
//                 return i
//             }
//             continue
//         }
//         if nums[i]>nums[i-1] && nums[i] > nums[i+1]{
//             return i
//         }
//     }
//     return -1
// }
// 利用每个元素都不相等
// func findPeakElement(nums []int) int {
//     for i:=0;i<len(nums)-1;i++{
//         if nums[i]>nums[i+1]{
//             return i
//         }
//     }
//     return len(nums)-1
// }
// 递归二分
func findPeakElement(nums []int) int {
	var search func([]int,int,int)int
	search = func(nums []int,left int,right int)int{
		if left==right {
			return left
		}
		mid:=right-(right-left+1)/2
		// 递减序列中 在左边找
		if nums[mid]>nums[mid+1]{
			return search(nums,left,mid)
		}
		// 递增序列中 在右边找
		return search(nums,mid+1,right)
	}
	return search(nums,0,len(nums)-1)
}


