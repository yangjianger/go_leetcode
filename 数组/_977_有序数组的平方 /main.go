package main

import "fmt"

/*
给你一个按 非递减顺序 排序的整数数组 nums，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。
示例 1：
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]

提示：
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 已按 非递减顺序 排序
 */
/*
解题思路：
	因为数组已经是排好序了的，在小于0的区间内，如果平方，那应该在一段区间内应该是递减的。在大于0区间内是递增的
 */
func main(){
	a := []int{-7,-3,2,3,11}
	fmt.Println(sortedSquares(a))
}

func sortedSquares(nums []int) []int {

	var (
		n = len(nums)
		lastIndex = -1
		res = make([]int, 0, n)
	)

	for i := 0; i < n && nums[i] < 0; i++ {
		lastIndex = i
	}

	for i, j := lastIndex, lastIndex + 1; i >= 0 || j < n;{
		if i < 0{
			res = append(res, nums[j] * nums[j])
			j++
		}else if j == n{
			res = append(res, nums[i] * nums[i])
			i--
		}else if nums[j] * nums[j] < nums[i] * nums[i]{
			res = append(res, nums[j] * nums[j])
			j++
		}else{
			res = append(res, nums[i] * nums[i])
			i --
		}
	}
	
	return res
}
