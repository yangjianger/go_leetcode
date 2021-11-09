/*
 * @Author: yangjiang
 * @Date: 2021-10-09 15:57:13
 * @Email: 1084914120@qq.com
 * @LastEditors: yangjiang
 * @LastEditTime: 2021-10-09 16:15:05
 * @Description: file information
 * @Company: jiguo
 */
package main

import "fmt"

//leetcode 53
func main() {
	nums := []int{-2, 8, -2, 1, 9}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	//状态转移方程
	//if dp[i-1] <= 0  { dp[i] = nums[i] }else{ dp[i] = dp[i-1] + nums[i] }
	//res = max(res, dp[i])
	//初始化
	p := nums[0]
	res := p
	for i := 1; i < len(nums); i++ {
		if p <= 0 {
			p = nums[i]
		} else {
			p = p + nums[i]
		}
		res = max(res, p)
	}

	return res
}

//因为只考虑前一个值，有点浪费空间
func maxSubArray1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	//状态转移方程
	//if dp[i-1] <= 0  { dp[i] = nums[i] }else{ dp[i] = dp[i-1] + nums[i] }
	//res = max(res, dp[i])
	//初始化
	dp[0] = nums[0]

	res := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
