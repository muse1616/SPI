package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	// dot split
	arr1, arr2 := strings.Split(version1, "."), strings.Split(version2, ".")
	i, j := 0, 0
	for i < len(arr1) || j < len(arr2) {
		if i < len(arr1) && j < len(arr2) {
			// 比较修订号 arr[i] arr[j]
			v1, _ := strconv.Atoi(arr1[i])
			v2, _ := strconv.Atoi(arr2[j])
			if v1 == v2 {
				i++
				j++
			} else if v1 > v2 {
				return 1
			} else if v1 < v2 {
				return -1
			}
		} else if i < len(arr1) {
			// 此时j已经无了 只有i后面有的不是非0直接返回1
			for i < len(arr1) {
				v1, _ := strconv.Atoi(arr1[i])
				if v1 != 0 {
					return 1
				}
				i++
			}
			return 0
		} else if j < len(arr2) {
			for j < len(arr2) {
				v2, _ := strconv.Atoi(arr2[j])
				if v2 != 0 {
					return -1
				}
				j++
			}
			return 0
		}
	}
	return 0
}
