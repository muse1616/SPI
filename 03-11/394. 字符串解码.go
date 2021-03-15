package main

import "strconv"

// 利用栈
// 1. 如果读到数字 解析出整个数字入栈
// 2. 如果是字母或者左括号 直接入栈
// 3. 如果读到右括号 栈pop直到左括号 pop出的字符组合成字符串 然后根据栈顶(此时栈顶一定是数字)
// 根据数字和字符串生成新的字符串 放入栈顶
// 4. 读完时栈中全部组合 即答案
// 数字ascii码 48 - 57
func decodeString(s string) string {
	// // 数组模拟栈
	stack := make([]string, 0)
	for i := 0; i < len(s); {
		if s[i] >= 48 && s[i] <= 57 {
			// 持续读直到数字
			start := i
			for s[i] >= 48 && s[i] <= 57 {
				i++
			}
			stack = append(stack, s[start:i])
		} else if s[i] == ']' {
			// 右括号时进行逻辑处理
			// 3. 如果读到右括号 栈pop直到左括号 pop出的字符组合成字符串 然后根据栈顶(此时              栈顶一定是数字)
			// 根据数字和字符串生成新的字符串 放入栈顶
			str := ""
			for stack[len(stack)-1] != "[" {
				// 组合 注意顺序
				str = stack[len(stack)-1] + str
				// pop
				stack = stack[:len(stack)-1]
			}
			// pop 左括号
			stack = stack[:len(stack)-1]
			// 此时栈顶是数量 根据数量生成字符串放入栈
			str = generate(str, stack[len(stack)-1])
			// pop 数字
			stack = stack[:len(stack)-1]
			// push新字符串
			stack = append(stack, str)
			// 指针后移
			i++
		} else {
			// 此时是字母或者左括号直接放入即可
			stack = append(stack, string(s[i]))
			i++
		}
	}
	ans := ""
	for _, v := range stack {
		ans += v
	}
	return ans
}
func generate(str string, numStr string) string {
	// 字符串转数字
	num, _ := strconv.Atoi(numStr)
	ans := ""
	for i := 0; i < num; i++ {
		ans += str
	}
	return ans
}
