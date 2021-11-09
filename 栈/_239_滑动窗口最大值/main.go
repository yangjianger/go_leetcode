package main

/*
给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k个数字。滑动窗口每次只向右移动一位。返回滑动窗口中的最大值。

解题思路：
1.这里需要用到双指针+双端队列.这里队列是逐渐递减的
	a,两个指针分别为l,r。 l = r - k + 1。
	b,创建返回数组，数组长度为 len(nums) - k + 1
	c,开始循环数组
		c1: 判断 nums[i] 是否大于队尾元素，如果大于，队尾元素出队。nums[i]入队。
		c2: 判断 w>=0时，判断队列队首元素是否无效，条件就是 队首元素是否 >= w, 如果无效出队。
		c3: 最后返回数组res[w] = 队首元素

 */

type Queue struct {
	eles []int
}

func initQueue(val int) *Queue{
	return &Queue{eles: make([]int, val, val)}
}

func (q *Queue) IsEmpty() bool{
	return len(q.eles) == 0
}

func (q *Queue) PushTail(val int){
	q.eles[len(q.eles)-1] = val
}

func (q *Queue) PopTail() int{
	val := q.eles[len(q.eles)-1]
	q.eles = q.eles[:len(q.eles)-1]
	return val
}

func (q *Queue) GetTail() int{
	return q.eles[len(q.eles)-1]
}

func (q *Queue) PopHead() int{
	val := q.eles[0]
	q.eles = q.eles[1:]
	return val
}

func (q *Queue) GetHead() int{
	return q.eles[0]
}


func maxSlidingWindow(nums []int, k int) []int {

	if nums == nil || len(nums) == 0 || len(nums) == 1 || k < 1{
		return nums
	}

	var (
		qCount = len(nums) - k + 1
		q = initQueue(len(nums))
		res = make([]int, qCount)
		r = 0
		l = -1
	)

	for ; r < len(nums); r++ {
		//判断队列是否顺序递增
		for !q.IsEmpty() && nums[r] >= nums[q.GetTail()]{
			q.PopTail()
		}
		q.PushTail(r)

		l = r - k + 1
		if l < 0{
			continue
		}

		if q.GetHead() < l{
			q.PopHead()
		}
		res[l] = nums[q.GetHead()]
	}	

	return res
}