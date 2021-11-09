/*
 * @Author: yangjiang
 * @Date: 2021-06-12 09:41:48
 * @Email: 1084914120@qq.com
 * @LastEditors: yangjiang
 * @LastEditTime: 2021-10-11 15:16:42
 * @Description: file information
 * @Company: jiguo
 */
package main

import "fmt"

func main() {
	fmt.Println(a(2))
}

func a(n int) []string {
	res := make([]string, 0)
	dfs(n, n, "", res)
	return res
}

func dfs(left, right int, curStr string, res []string) {

	if left == 0 && right == 0 { // 左右括号都不剩余了，递归终止
		res = append(res, curStr)
		return
	}

	if left > 0 { // 如果左括号还剩余的话，可以拼接左括号
		dfs(left-1, right, curStr+"(", res)
	}
	if right > left { // 如果右括号剩余多于左括号剩余的话，可以拼接右括号
		dfs(left, right-1, curStr+")", res)
	}

}
