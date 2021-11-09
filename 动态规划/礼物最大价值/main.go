package main

import (
	"fmt"
	"math"
)

//计算最大值
func maxValue(grid [][]int) int {

	if len(grid) == 0{
		return 0
	}

	var (
		cols = len(grid[0])
		rows = len(grid)
		dp = make([][]int, rows)
	)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, cols)
	}

	dp[0][0] = grid[0][0]
	for col := 1; col < cols; col++ {
		dp[0][col] =  grid[0][col] + dp[0][col-1]
	}

	for row := 1; row < rows; row++ {
		dp[row][0] =  grid[row][0]  + dp[row-1][0]
	}

	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			dp[row][col] = int(math.Max(float64(dp[row-1][col]),float64(dp[row][col-1]))) + grid[row][col]
		}
	}

	return dp[rows-1][cols-1]
}
func main(){

	arr := [][]int{
		{1, 2, 3},
		{4, 5 ,6},
		{7, 8, 9},
	}

	fmt.Println(maxValue(arr))
}
