package main

/*
给定一个整数数组，编写一个函数，找出索引m和n，只要将索引区间[m,n]的元素排好序，整个数组就是有序的。注意：n-m尽量最小，也就是说，找出符合条件的最短序列。函数返回值为[m,n]，
若不存在这样的m和n（例如整个数组是有序的），请返回[-1,-1]。
示例：

输入： [1,2,4,7,10,11,7,12,6,7,16,18,19]
输出： [3,9]
提示：

0 <= len(array) <= 1000000
 */

/*
解题思路：
	1.这种是寻找逆序对
	2.从左到右应该是递增的，所以这里需要存储一个当前扫描的最大值，如果后面的值小于最大值，那存储。如果大于最大值，
      将最大值重新赋值,因为只寻找最小区间，所以相等的时候，不考虑
 */
func main(){

}

func subSort(array []int) []int {

	if len(array) == 0{
		return []int{-1, -1}
	}

	var (
		lens = len(array)
		max = array[0]
		l = -1
		min = array[lens - 1]
		r = -1
	)

	for i := 1; i < lens; i++ {
		v := array[i]
		if v >= max{
			max = v
		}else{
			//小于了最大值
			r = i
		}
	}

	if r == -1{
		return []int{-1, -1}
	}

	//从右到左，逐渐变小
	for i := lens - 2; i >= 0; i-- {
		v := array[i]
		if v <= min{
			min = v
		}else{
			//大于了最小值
			l = i
		}
	}

	return []int{l, r}
}
