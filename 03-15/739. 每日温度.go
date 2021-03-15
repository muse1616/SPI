package main

// 单调递减栈
func dailyTemperatures(T []int) (res []int) {
	res = make([]int, len(T))
	// stack中放index即可
	stack := make([]int, 0)
	for i := 0; i < len(T); i++ {
		// 此时遇到了递增
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	// 此时栈中所有序号都为0
	for _, v := range stack {
		res[v] = 0
	}

	return res
}
