package main

func maximumSwap(num int) int {
	str := []byte(strconv.Itoa(num))
	indexes := make([]int, 10)
	//记录每个元素的索引值，如果又重复，记录最右边的
	for i, b := range str {
		indexes[b-'0'] = i
	}
	//从左到右遍历str，然后每个元素和9-0对比，
	//例如：如果9存在，并且9的索引比当前大（在右边），那么换位置，返回结果
	for i, b := range str {
		for d := 9; d > int(b-'0'); d-- {
			if indexes[d] > i {
				str[i], str[indexes[d]] = str[indexes[d]], str[i]
				res, _ := strconv.Atoi(string(str))
				return res
			}
		}
	}
	return num
}

// // 选择一个最大的数字 交换前面一个最小的数字
// // 所有数字递减的时候 不需要交换
// // n2复杂度
// func maximumSwap(num int) int {
//     // 转成数组
//     arr:=toArr(num)
//     for i:=0;i<len(arr);i++{
//         maxIndex:=i
//         for j:=len(arr)-1;j>=i+1;j--{
//             // 在 j..最后找一个最大的数字 需要记录序号
//             // 此时需要反序找
//             if arr[j]>arr[maxIndex]{
//                 maxIndex = j
//             }
//         }
//         // 此时需要交换
//         if maxIndex!=i{
//             arr[i],arr[maxIndex] = arr[maxIndex],arr[i]
//             break
//         }
//     }

//     return toNum(arr)
// }
// func toArr(num int)(ans []int){
//     for num!=0{
//         ans = append([]int{num%10},ans...)
//         num/=10
//     }
//     return
// }
// func toNum(arr []int)(num int){
//     for i:=0;i<len(arr);i++{
//         num=num*10+arr[i]
//     }
//     return
// }
