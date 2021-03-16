package main
func isPalindrome(s string) bool {
	// 双指针
	left,right:=0,len(s)-1
	for left<right{
		// 遇到除了字母和数字的直接跳过
		for left<right && !isLetterOrDigit(s[left]){
			left++
		}
		for left<right && !isLetterOrDigit(s[right]){
			right--
		}
		// 判断相同
		if left==right{
			return true
		}
		if !check(s[left],s[right]){
			return false
		}
		left++
		right--
	}
	// 此时包括了空字符串
	return true
}
// 此时需要忽略大小写
func check(a,b byte)bool{
	// 换成大写即可
	if a>=97{
		a-=32
	}
	if b>=97{
		b-=32
	}
	if a==b {
		return true
	}
	return false
}
func isLetterOrDigit(a byte)bool{
	// a-z：97-122，A-Z：65-90，0-9：48-57。
	if (a >= 97 && a<=122) || (a>=65 &&a<=90) || (a>=48&&a<=57){
		return true
	}
	return false
}