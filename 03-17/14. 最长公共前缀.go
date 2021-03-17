package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	index := 0
	for {
		if index >= len(strs[0]) {
			break
		}
		// 取第一个字符串的字符
		letter := strs[0][index]
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i]) || strs[i][index] != letter {
				return strs[0][:index]
			}
		}
		index++
	}

	return strs[0][:index]
}
