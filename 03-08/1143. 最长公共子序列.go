package main

func longestCommonSubsequence(text1 string, text2 string) (ans int) {
	// no bound

	// 2 - dp
	dp:=make([][]int,len(text1)+1)
	for i:=0;i<len(dp);i++{
		dp[i] = make([]int,len(text2)+1)
	}

	// dp[i][j] text1前i个. text2前j个的最长公共子序列
	for i:=1;i<=len(text1);i++{
		for j:=1;j<=len(text2);j++{
			// 若第i个字母等于第j个字母 当前dp为前一个dp加上1（当前字母相同所以加1）
			if text1[i-1] == text2[j-1]{
				dp[i][j] = dp[i-1][j-1]+1
			}else{
				// 若果当前字母不同
				// 比如 ...a 和 ...b 这两个字母不同
				// 第一个字符串长度为i 第二个字符串长度为j
				// 此时的dp 就等于 第一个字符串把a去掉 即...和...b的最大长度 即 dp[i-1][j]
				// 或者第二个字符串把b去掉 即...a和...b的最大长度 即dp[i][j-1]
				// 这两个值取一个 这时不能两个都去掉 即不能等于dp[i-1][j-1]
				// 解释: 如果把 a b都去掉 若第一个字符串是...ba 第二个字符串是...b 此时去掉了a 并没有对当前的dp产生影响
				dp[i][j] = max(dp[i-1][j],dp[i][j-1])
			}
		}
	}
	// 最右下角那项
	return dp[len(text1)][len(text2)]
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}