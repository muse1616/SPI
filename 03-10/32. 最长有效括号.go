package main

func longestValidParentheses(s string) int {
	ans := 0
	//dp
	dp := make([]int, len(s))
	// 因为已左括号结尾的串dp必定为0 所以只需要判断右括号时的情况
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				// 1. ...()
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					// 2. ()
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				// 此时s[i-1] == ')' 字符串形如 ...)) 需要前面有一个左括号
				// 即 i-dp[i-1]>=0 同时前方出现一个( 形如 ( ... ))
				// 1. ...( ... ))
				if i-dp[i-1] >= 2 {
					dp[i] = 2 + dp[i-1] + dp[i-dp[i-1]-2]
				} else {
					// 2. (...))
					dp[i] = dp[i-1] + 2
				}
			}
			ans = max(ans, dp[i])
		}

	}

	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
