/*
 * @Author: yangjiang
 * @Date: 2021-10-09 16:27:57
 * @Email: 1084914120@qq.com
 * @LastEditors: yangjiang
 * @LastEditTime: 2021-10-21 15:20:15
 * @Description: file information
 * @Company: jiguo
 */
package main

import "fmt"

func main() {
	nums := []int{7, 7, 7, 7, 7, 7, 7}
	fmt.Println(lengthOfLIS(nums))
}

//状态转移方程：dp[i] = max(dp[i], dp[j]+1) _max = max(dp[i], _max)
func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	_max := 1
	for i := 0; i < len(nums); i++ {
		//状态的初始值，所有的dp[i]初始化都是1
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			dp[i] = max(dp[i], dp[j]+1)
		}
		_max = max(dp[i], _max)
	}
	return _max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
