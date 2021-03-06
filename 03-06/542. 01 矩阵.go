package main

import "math"

func updateMatrix(matrix [][]int) [][]int {
	// bound
	if len(matrix) == 0{
		return [][]int{}
	}
	// 2-dp
	dp:=make([][]int,len(matrix))
	for i:=0;i<len(dp);i++{
		dp[i] = make([]int,len(matrix[0]))
		// default
		for j:=0;j<len(dp[i]);j++{
			if matrix[i][j] == 0{
				dp[i][j] = 0
			}else{
				// 此处不能等于 len*len
				dp[i][j] = math.MaxInt32
			}
		}
	}
	// [0,0] -> [-1,-1]
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[i]);j++{
			//left and up
			if i>=1{
				dp[i][j] = min(dp[i][j],dp[i-1][j]+1)
			}
			if j>=1{
				dp[i][j] = min(dp[i][j],dp[i][j-1]+1)
			}
		}
	}
	// [-1,-1] -> [0,0]
	for i:=len(matrix)-1;i>=0;i--{
		for j:=len(matrix[0])-1;j>=0;j--{
			// right and bottom
			if i<len(matrix)-1{
				dp[i][j] = min(dp[i][j],dp[i+1][j]+1)
			}
			if j<len(matrix[0])-1{
				dp[i][j] = min(dp[i][j],dp[i][j+1]+1)
			}
		}
	}
	return dp
}
func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}