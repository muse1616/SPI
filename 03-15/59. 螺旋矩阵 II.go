package main

func generateMatrix(n int) (ans [][]int) {
	ans = make([][]int, n)
	for i := 0; i < len(ans); i++ {
		ans[i] = make([]int, n)
	}
	// 四个边界
	left, right, up, bottom := 0, n-1, 0, n-1
	start, end := 1, n*n
	for start <= end {
		// 左到右 上+1
		for i := left; i <= right; i++ {
			ans[up][i] = start
			start++
		}
		up++
		// 上到下
		for i := up; i <= bottom; i++ {
			ans[i][right] = start
			start++
		}
		right--
		// 右到左
		for i := right; i >= left; i-- {
			ans[bottom][i] = start
			start++
		}
		bottom--
		// 下到上
		for i := bottom; i >= up; i-- {
			ans[i][left] = start
			start++
		}
		left++
	}
	return
}
