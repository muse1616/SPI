package main

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
			continue
		}
		// 此时 s[left]!=s[right] 需要确定删除哪个
		// 1.删除左边
		nL, nR := left+1, right
		for nL < nR {
			if s[nL] != s[nR] {
				break
			}
			nL++
			nR--
		}
		if nL >= nR {
			return true
		}

		// 2.删除右边
		nL, nR = left, right-1
		for nL < nR {
			if s[nL] != s[nR] {
				break
			}
			nL++
			nR--
		}
		if nL >= nR {
			return true
		}
		// 3.两种都尝试后依然不可以 即返回false
		return false
	}

	return true
}
