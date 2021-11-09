/*
 * @Author: yangjiang
 * @Date: 2021-10-14 16:02:11
 * @Email: 1084914120@qq.com
 * @LastEditors: yangjiang
 * @LastEditTime: 2021-10-14 16:35:42
 * @Description: file information
 * @Company: jiguo
 */
package main

import "fmt"

/**
初始化： dp[i, 0] = dp[0, j] = 0
状态转移方程：if(str1[i-1] == str2[j-1]) dp[i, j] = dp[i-1, j-1] + 1 else dp[i, j] = 0
最后：所有dp[i,  j] 中的最大值
 * @description:
 * @param {*}
 * @return {*}
*/
func main() {
	fmt.Println(lcs("abcde", "ace"))
}

func lcs(text1, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}

	_max := 0

	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] != text2[j-1] {
				continue
			}
			dp[i][j] = dp[i-1][j-1] + 1
			_max = max(dp[i][j], _max)
		}
	}

	return _max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
