package main

//给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
func integerBreak(n int) int {

	mem := make([]int, n+1)
	return breakInteger(n, mem)
}

func breakInteger(n int, mem []int) int{

	if n == 1{
		return 1
	}

	if mem[n] != 0{
		return mem[n]
	}

	res := 0

	for i := 1; i < n; i++{
		res = max3(res, i * (n - i),  i * breakInteger(n-i, mem))
	}
	mem[n] = res

	return res
}

func max(a, b int) int{
	if a > b{
		return a
	}

	return b
}

func max3(a, b, c int) int{
	return max(a, max(b, c))
}
