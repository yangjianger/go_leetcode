/*
 * @Author: yangjiang
 * @Date: 2021-10-14 19:08:04
 * @Email: 1084914120@qq.com
 * @LastEditors: yangjiang
 * @LastEditTime: 2021-10-21 11:16:46
 * @Description: file information
 * @Company: jiguo
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	values := []int{6, 3, 5, 5, 6}
	weights := []int{2, 2, 6, 5, 4}
	capacity := 10
	fmt.Println(maxValue3(values, weights, capacity))
}

func maxValue(values, weights []int, capacity int) int {
	if values == nil || weights == nil {
		return 0
	}
	if len(values) == 0 || len(weights) == 0 || capacity <= 0 {
		return 0
	}
	if len(values) != len(weights) {
		return 0
	}

	dp := make([][]int, len(values)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, capacity+1)
	}
	for i := 1; i <= len(values); i++ {
		for j := 1; j <= capacity; j++ {
			//状态转移方程： 最后一个不选，如果选了，就去一个最大的
			if j < weights[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i-1]]+values[i-1])
			}
		}
	}
	return dp[len(values)][capacity]
}

func maxValue2(values, weights []int, capacity int) int {
	if values == nil || weights == nil {
		return 0
	}
	if len(values) == 0 || len(weights) == 0 || capacity <= 0 {
		return 0
	}
	if len(values) != len(weights) {
		return 0
	}

	dp := make([]int, capacity+1)

	for i := 1; i <= len(values); i++ {
		for j := capacity; j >= 1; j-- {
			//状态转移方程： 最后一个不选，如果选了，就去一个最大的
			if j >= weights[i-1] {
				dp[j] = max(dp[j], dp[j-weights[i-1]]+values[i-1])
			}
		}
	}
	return dp[capacity]
}

func maxValue3(values, weights []int, capacity int) int {
	if values == nil || weights == nil {
		return 0
	}
	if len(values) == 0 || len(weights) == 0 || capacity <= 0 {
		return 0
	}
	if len(values) != len(weights) {
		return 0
	}

	dp := make([]int, capacity+1)

	for i := 1; i <= len(values); i++ {
		for j := capacity; j >= weights[i-1]; j-- {
			//状态转移方程： 最后一个不选，如果选了，就去一个最大的
			dp[j] = max(dp[j], dp[j-weights[i-1]]+values[i-1])
		}
	}
	return dp[capacity]
}

//恰好装满
func maxValueExactly(values, weights []int, capacity int) int {
	if values == nil || weights == nil {
		return 0
	}
	if len(values) == 0 || len(weights) == 0 || capacity <= 0 {
		return 0
	}
	if len(values) != len(weights) {
		return 0
	}

	dp := make([]int, capacity+1)
	for i := 1; i <= capacity; i++ {
		dp[i] = math.MinInt
	}

	for i := 1; i <= len(values); i++ {
		for j := capacity; j >= weights[i-1]; j-- {
			//状态转移方程： 最后一个不选，如果选了，就去一个最大的
			dp[j] = max(dp[j], dp[j-weights[i-1]]+values[i-1])
		}
	}

	if dp[capacity] < 0 {
		return -1
	}
	return dp[capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
