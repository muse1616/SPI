package main
func longestPalindrome(s string) string {
	dp:=make([][]bool,len(s))
	for i:=0;i<len(dp);i++{
		dp[i] = make([]bool,len(s))
	}
	ans:=""
	for length:=0;length<len(s);length++{
		for left:=0;left+length<len(s);left++{
			right:=left+length
			if length == 0{
				dp[left][right] = true
			}else if length == 1{
				if s[left]==s[right]{
					dp[left][right]=true
				}
			}else{
				if s[left]==s[right]{
					dp[left][right] = dp[left+1][right-1]
				}
			}
			if dp[left][right] == true && length+1>len(ans){
				ans = s[left:right+1]
			}
		}
	}


	return ans
}