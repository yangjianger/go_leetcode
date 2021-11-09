package main

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
/**
解题思路：
	1. 这个其实就是斐波拉切思路，比如，第一步爬了一阶，那么后面就还有n-1阶， 第一步爬了两阶，后面就还有n-2阶， 总的就是 f(n) = f(n-1) + f(n-2)

 */
func climbStairs(n int) int {

	first := 1
	second := 2
	for i := 2; i < n; i++ {
		second = first + second
		first = second - first
	}
	return second
}
