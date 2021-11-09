package main

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//示例 1： 输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]

//示例 2：
//
//输入：n = 1
//输出：["()"]
func generateParenthesis(n int) []string {

	res := make([]string, 0, n * 2)
	dfs(0, 0, n, "", &res)

	return res
}

//当字符串长度 == n * 2 的时候，结束循环
//当左括号<n的时候， 左括号+1
//当右括号小于左括号的时候，右括号+1
func dfs(left, right, n int, s string, res *[]string){

	if len(s) == n * 2{
		*res = append(*res, s)
		return
	}

	if left < n{
		dfs(left + 1, right, n, s + "(", res)
	}

	if right < left{
		dfs(left, right + 1, n, s + ")", res)
	}

}