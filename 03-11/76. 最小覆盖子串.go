package main

// hash+sliding window
func minWindow(s string, t string) string {
	if s == t {
		return s
	}
	// 长度不符
	if len(s) < len(t) {
		return ""
	}
	ans := ""
	// 字母及出现次数
	dic := make(map[byte]int, 0)
	for i := 0; i < len(t); i++ {
		dic[t[i]]++
	}
	// 滑动窗口的hash表 每一个字母都不少于即可
	slideHash := make(map[byte]int, 0)
	// 滑动窗口左右边界
	left, right := 0, 0
	slideHash[s[right]] = 1
	for right < len(s) {
		// 滑动窗口中包含了字符串
		if isContain(dic, slideHash) {
			// 一个字符 绝对是最短
			if left == right {
				ans = string(s[left])
				return ans
			}
			//左窗口后移
			for isContain(dic, slideHash) && left <= right {
				// hash中删除
				slideHash[s[left]]--
				left++
			}
			if ans == "" || len(s[left-1:right+1]) < len(ans) {
				ans = s[left-1 : right+1]
			}
		} else {
			// 滑动窗口中不包含 右窗口向后
			right++
			// 退出条件
			if right == len(s) {
				break
			}
			slideHash[s[right]]++
		}
	}
	return ans
}

// 窗口中是否包含hash
func isContain(dic map[byte]int, slideHash map[byte]int) bool {
	for k, v := range dic {
		if slideHash[k] < v {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
