/*
 * @Author: yangjiang
 * @Date: 2021-10-09 17:54:11
 * @Email: 1084914120@qq.com
 * @LastEditors: yangjiang
 * @LastEditTime: 2021-10-21 11:43:55
 * @Description: file information
 */
package main

import "fmt"

/*
	1.定义状态
		dp[i, j]
		i ∈ [0, nums1.length]
		j ∈ [0, nums2.length]
	2.状态初始化
		dp[i, 0]  = dp[j, 0] = 0
	3.状态转移方程
		当num1[i-1] == num2[j-1] 那么就是 dp[i-1, j-1] + 1
		当num1[i-1] != num2[j-1] 那么就是 max(dp[i, j-1], dp[i-1, j])
*/

func main() {
	num1 := "abc"
	num2 := "bcd"
	fmt.Println(lcs3(num1, num2))
}

func lcs3(num1, num2 string) int {
	if len(num1) == 0 || len(num2) == 0 {
		return 0
	}

	dp := make([][]int, 2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(num2)+1)
	}

	row := 0
	prevRow := 0
	for i := 1; i <= len(num1); i++ {
		for j := 1; j <= len(num2); j++ {
			row = i % 2
			prevRow = (i - 1) % 2
			if num1[i-1] == num2[j-1] {
				dp[row][j] = dp[prevRow][j-1] + 1
			} else {
				dp[row][j] = max(dp[row][j-1], dp[prevRow][j])
			}
		}
	}

	return dp[len(num1)%2][len(num2)]

}

func lcs2(num1, num2 string) int {
	return _lcs2(num1, len(num1), num2, len(num2))
}
func _lcs2(num1 string, i int, num2 string, j int) int {
	if i == 0 || j == 0 {
		return 0
	}

	if num1[i-1] == num2[j-1] {
		return _lcs2(num1, i-1, num2, j-1) + 1
	}
	//最后一个元素不相等
	return max(
		_lcs2(num1, i, num2, j-1),
		_lcs2(num1, i-1, num2, j))
}

func lcs(num1, num2 []int) int {
	return _lcs(num1, len(num1), num2, len(num2))
}

/**
 * @description: 求num1前i个元素和num2前j个元素最长公共子序列
 * @param {[]int} num1
 * @param {int} i
 * @param {[]int} num2
 * @param {int} j
 * @return {*}
 */
func _lcs(num1 []int, i int, num2 []int, j int) int {
	if i == 0 || j == 0 {
		return 0
	}

	if num1[i-1] == num2[j-1] {
		return _lcs(num1, i-1, num2, j-1) + 1
	}
	//最后一个元素不相等
	return max(
		_lcs(num1, i, num2, j-1),
		_lcs(num1, i-1, num2, j))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
