package main

type pair struct{ x, y int }

// 上下左右
var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	// 访问数组
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	var dfs func(int, int, int) bool
	dfs = func(i, j, index int) bool {
		if board[i][j] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		// 已访问
		visited[i][j] = true
		// 最后还原访问数组
		defer func() {
			visited[i][j] = false
		}()
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < len(board) && 0 <= newJ && newJ < len(board[0]) && !visited[newI][newJ] {
				if dfs(newI, newJ, index+1) {
					return true
				}
			}
		}
		return false
	}
	// 遍历
	for i, row := range board {
		for j := range row {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
