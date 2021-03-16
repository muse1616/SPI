package main

func generateParenthesis(n int) (ans []string) {
	var backtrace func(int, int, string)
	backtrace = func(left int, right int, current string) {
		if left+right == 2*n {
			if left == right {
				ans = append(ans, current)
			}
			return
		}
		// 放左或者放右
		if left == n {
			backtrace(left, right+1, current+")")
			return
		}
		// 左大于右才能放右
		if left > right {
			backtrace(left, right+1, current+")")
		}
		// 放左
		backtrace(left+1, right, current+"(")
	}
	backtrace(0, 0, "")
	return
}
