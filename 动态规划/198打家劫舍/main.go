package main

//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

func rob(nums []int) int {
	mem := make([]int, len(nums))

	return tryRob(nums, 0, mem)
}

func max(a, b int) int{
	if a > b{
		return a
	}

	return b
}

func tryRob(nums []int, index int, mem []int) int{

	if index >= len(nums){
		return 0
	}

	if mem[index] != 0{
		return mem[index]
	}

	res := 0
	for i := index; i < len(nums);  i++{
		res = max(res, nums[i] + tryRob(nums, i + 2, mem))
	}

	mem[index] = res
	return res
}