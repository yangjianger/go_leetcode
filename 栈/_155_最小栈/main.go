package main

import "math"

/*
题目描述：
	设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
解题思路：
	这里需要创建一个最小栈，每次入栈的时候，需要将最小栈的栈顶数据和当前入栈元素比较，最小栈入栈最小元素

 */

type MinStack struct {
	stack []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {

	return MinStack{
		stack: []int{},
		minStack: []int{math.MaxInt64},
	}
}

func min(x, y int) int{
	if x < y{
		return x
	}
	return y
}

func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack, val)
	top := this.minStack[len(this.minStack) - 1]
	this.minStack = append(this.minStack, min(top, val))
}

func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack) - 1]
}
