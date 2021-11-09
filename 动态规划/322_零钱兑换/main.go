/*
 * @Author: yangjiang
 * @Date: 2021-10-08 10:21:12
 * @Email: 1084914120@qq.com
 * @LastEditors: yangjiang
 * @LastEditTime: 2021-10-09 14:58:35
 * @Description: file information
 * @Company: jiguo
 */
package main

import (
	"fmt"
	"math"
)

/*
解题思路：
	1.考虑第一次选25或者20或者5或者1之后的最小值 + 1
	2.考虑递归基：当n=25或者20或者5或者1时，最优化的解应该是1
	3.当n<=0的时候，需要返回最大值，如果是负数或者0的时候，min就会报错
*/
func main() {

	//fmt.Println(coins3(59))
	faces := []int{1, 5, 20, 25}
	fmt.Println(coins4(8, faces))
}

/**
 * @Author: yangjiang
 * @Date: 2021-10-09 14:30:25
 * @Title: 动态传入faces
 * @description:
 * @param {int} n
 * @param {[]int} faces
 * @return {*}
 */
func coins4(n int, faces []int) int {
	if n < 1 || faces == nil || len(faces) == 0 {
		return -1
	}

	dp := make([]int, n+1)
	mins := 0
	for i := 1; i <= n; i++ {
		mins = math.MaxInt
		for j := 0; j < len(faces); j++ {
			if i < faces[j] || dp[i-faces[j]] < 0 {
				continue
			}
			mins = min(dp[i-faces[j]], mins)
		}
		if mins == math.MaxInt {
			dp[i] = -1
		} else {
			dp[i] = mins + 1
		}

	}

	return dp[n]
}

/**
 * @Author: yangjiang
 * @Date: 2021-10-09 10:54:41
 * @Title: 递推，自底向上
 * @description:
 * @param {int} n
 * @return {*}
 */
func coins3(n int) int {
	if n < 1 {
		return -1
	}
	dp := make([]int, n+1)
	faces := make([]int, len(dp))
	_min := 0

	for i := 1; i <= n; i++ {
		_min = math.MaxInt
		if i >= 1 && dp[i-1] < _min {
			_min = dp[i-1]
			faces[i] = 1
		}
		//_min = dp[i-1]
		if i >= 5 && dp[i-5] < _min {
			_min = dp[i-5]
			faces[i] = 5
		}
		if i >= 20 && dp[i-20] < _min {
			_min = dp[i-20]
			faces[i] = 20
		}
		if i >= 25 && dp[i-25] < _min {
			_min = dp[i-25]
			faces[i] = 25
		}
		dp[i] = _min + 1

	}
	printN(n, faces)
	return dp[n]
}

/**
 * @Author: yangjiang
 * @Date: 2021-10-09 14:16:21
 * @Title: 打印faces
 * @description:
 * @param {int} n
 * @param {[]int} faces
 * @return {*}
 */
func printN(n int, faces []int) {
	for n > 0 {
		fmt.Printf("%d\t", faces[n])
		//这里是取最后一次
		n -= faces[n]
	}
	fmt.Println()
}

/**
 * @Author: yangjiang
 * @Date: 2021-10-09 10:19:16
 * @Title: 记忆化搜索，自顶向下
 * @description:
 * @param {int} n
 * @return {*}
 */
func coins2(n int) int {
	if n < 1 {
		return -1
	}

	lens := n

	if lens < 25 {
		lens = 25
	}

	dp := make([]int, lens+1)
	dp[1] = 1
	dp[5] = 1
	dp[20] = 1
	dp[25] = 1

	return _coins2(n, dp)
}

func _coins2(n int, dp []int) int {
	if n < 1 {
		return math.MaxInt
	}

	if dp[n] != 0 {
		return dp[n]
	}

	min1 := min(_coins2(n-25, dp), _coins2(n-20, dp))
	min2 := min(_coins2(n-5, dp), _coins2(n-1, dp))
	dp[n] = min(min1, min2) + 1
	return dp[n]
}

/**
 * @Author: yangjiang
 * @Date: 2021-10-08 18:59:30
 * @description: 暴力递归，出现重叠子问题
 * @param {int} n
 * @return {*}
 */
func coins1(n int) int {
	if n <= 0 {
		return math.MaxInt
	}

	if n == 25 || n == 20 || n == 5 || n == 1 {
		return 1
	}

	min1 := min(coins1(n-25), coins1(n-20))
	min2 := min(coins1(n-5), coins1(n-1))

	return min(min1, min2) + 1
}

/**
 * @Author: yangjiang
 * @Date: 2021-10-09 10:15:55
 * @description: 最小值
 * @param {*} a
 * @param {int} b
 * @return {*}
 */
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
