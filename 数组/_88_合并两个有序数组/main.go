package main

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
初始化 nums1 和 nums2 的元素数量分别为  m 和 n 。你可以假设  nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]

输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]

*/
/*
解题思路：
	0.这里需要从最后一个元素开始比较，如果从第一个开始比较，多余的数据就不知道存储在哪里了
	1.需要三个指针，分别为：indexI1(num1最后一个元素),indexI2(nums2最后一个元素),curIndex(当前元素下标)
	2.临界条件，indexI2>=0,这里只考虑nums2，因为，nums1已经是排好序了的
	3.当indexI1>=0 && nums1[indexI1] >= nums2[indexI2],这里需要考虑，nums1[curIndex] = nums1[indexI1]
*/

func merge(nums1 []int, m int, nums2 []int, n int)  {
	var (
		indexI1 = m - 1
		indexI2 = n - 1
		curIndex = len(nums1) - 1
	)

	for  indexI2 >= 0{
		if indexI1 >= 0 && nums1[indexI1] >= nums2[indexI2]{
			nums1[curIndex] = nums1[indexI1]
			indexI1--
		}else{
			//indexI1 < 0 nums1[indexI1] < nums2[indexI2]
			nums1[curIndex] = nums2[indexI2]
			indexI2--
		}
		curIndex--
	}
}


