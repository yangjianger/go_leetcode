package main

import "fmt"

/*
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
此题中，我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。
示例 1：
输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]

示例 2：
输入：nums = [2,0,1]
输出：[0,1,2]

示例 3：
输入：nums = [0]
输出：[0]

示例 4：
输入：nums = [1]
输出：[1]

提示：
n == nums.length
1 <= n <= 300
nums[i] 为 0、1 或 2
*/

/*
解题思路：
	1.这里需要三个指针，分别代表 0(l)， 2(r) 和当前指针(cur)
	2.临界条件，cur > r, 这里不能等于，如果最后一次交换的值是0，2.这就有问题了
	3.判断当前指针如果为0，交换cur和l下标值，l和cur值++,
	4.如果为2，交换cur和r下标值，r--, cur不变
	5.如果为1，cur++
*/

func main(){
	arr :=  []int{2,0,1}
	sortColors(arr)

	fmt.Println(arr)
}

func sortColors(nums []int)  {
	var (
		l  = 0
		r = len(nums) - 1
		cur = 0
	)

	for  cur <= r{
		if nums[cur] == 0{
			nums[cur], nums[l] = nums[l], nums[cur]
			cur ++
			l ++
		}else if nums[cur] == 2{
			nums[cur], nums[r] = nums[r], nums[cur]
			r--
		}else{
			cur ++
		}
	}

}