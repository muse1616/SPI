package main

// 单调递减栈
func trap(height []int) (sum int) {
	if len(height) == 0 {
		return
	}
	// 序号栈
	stack := make([]int, 0)
	for i := 0; i < len(height); i++ {
		// 遇到了递增的
		for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
			// 出栈
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 弹出后栈空 说明 两数递增 无意义 跳出循环
			if len(stack) == 0 {
				break
			}
			// 计算宽度 这里计算的其实是宽度 因为是横着看的
			// 比如 2 1 2 此时横着计算 1*1 第一个1值得是2-1
			w := min(height[i], height[stack[len(stack)-1]]) - height[pop]
			sum = sum + (i-1-stack[len(stack)-1])*w
		}
		// 入栈
		stack = append(stack, i)
	}

	return

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
