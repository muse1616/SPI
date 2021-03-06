package main

// 深度搜索
func numIslands(grid [][]byte)(cnt int){
	var dfs func([][]byte,int,int)

	dfs = func(grid [][]byte,i int,j int) {
		if i<0 || i>=len(grid) || j<0 || j>=len(grid[i])|| grid[i][j] == '0'{
			return
		}
		grid[i][j] = '0'
		dfs(grid,i+1,j)
		dfs(grid,i-1,j)
		dfs(grid,i,j+1)
		dfs(grid,i,j-1)

	}


	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]);j++{
			if grid[i][j] == '1'{
				cnt++
				dfs(grid,i,j)
			}
		}
	}

	return
}
