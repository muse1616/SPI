package main

import (
	"sort"
	"strconv"
	"strings"
)

// 根据字典序排序
func largestNumber(nums []int) string {
	toStringArr := make([]string, len(nums))
	for i, num := range nums {
		toStringArr[i] = strconv.Itoa(num)
	}
	// 自定义排序
	sort.Slice(toStringArr, func(i, j int) bool {
		// 字典序排列
		return toStringArr[i]+toStringArr[j] >= toStringArr[j]+toStringArr[i]
	})
	ans := strings.Join(toStringArr, "")
	if ans[0] == '0' {
		return "0"
	}
	return ans
}
