package main

import (
	"fmt"
	"math"
)

/**
解题思路：分治思想，其实就是求 [begin, mid), [mid, end), 中的最大连续子串。有三种情况，可能最大就在左边，右边，还有一种就是左边，右边都有
	[begin, mid)	[mid, end)
	  [left,mid) [mid, right)
*/

func main() {
	nums := []int{2, 4}

	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	return _maxSubArray(nums, 0, len(nums))
}

func _maxSubArray(nums []int, begin, end int) int {
	//设置递归基
	if end-begin < 2 {
		return nums[begin]
	}

	mid := (end + begin) >> 1
	leftMax := math.MinInt
	leftSum := 0
	for i := mid - 1; i >= begin; i-- {
		leftSum += nums[i]
		leftMax = max(leftSum, leftMax)
	}

	rightMax := math.MinInt
	rightSum := 0
	for i := mid; i < end; i++ {
		rightSum += nums[i]
		rightMax = max(rightSum, rightMax)
	}

	return max(leftMax+rightMax, max(_maxSubArray(nums, begin, mid), _maxSubArray(nums, mid, end)))
}

func max(i, j int) int {
	return int(math.Max(float64(i), float64(j)))
}
