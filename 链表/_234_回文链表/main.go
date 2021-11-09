package main

//请判断一个链表是否为回文链表。
//解题思路
/*
1.将元素个数<3 和 >=3分开判断
	a. < 3, 如果元素个数为0，或者 1 返回true, 为2的时候，直接判断是否相等
	b. >= 3 , 获取中间节点(利用快慢指针，快指正.Next.Next==nil || .Next==nil)为循环界限，最后返回慢指针，就是中间节点。
		以中间节点.Next为根，反转链表。之后以 rHead != nil 为界，判断是否相等
 */

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {

	//head = nil  head = 1  head = 1->2
	if head == nil || head.Next == nil{
		return true
	}

	if head.Next.Next == nil{
		return head.Next.Val == head.Val
	}

	//获取中间节点
	middleNode := getMiddleNode(head)

	//反转链表
	rHead := reserv(middleNode.Next)

	for rHead != nil{
		if rHead.Val != head.Val{
			return false
		}

		head = head.Next
		rHead = rHead.Next
	}

	return true
}

//获取中间节点 右半部分链表节点的前一个节点
//比如 1->2->3->2->1 中的3就是中间节点
//1->2->2->1 中左边第一个2就是中间节点
//这里通过快慢指针找到
func getMiddleNode(head *ListNode) *ListNode{
	var (
		fastNode = head
		slowNode = head
	)

	for fastNode.Next != nil && fastNode.Next.Next != nil{
		fastNode = fastNode.Next.Next
		slowNode = slowNode.Next
	}

	return slowNode
}

//反转链表
func reserv(head *ListNode) *ListNode{

	var (
		newHead *ListNode
		tmp *ListNode
	)

	for head != nil{
		tmp = head.Next

		head.Next = newHead
		newHead = head

		head = tmp
	}

	return newHead
}

func main(){
	a := &ListNode{Next: nil, Val: 1}
	b := &ListNode{Next: nil, Val: 2}
	c := &ListNode{Next: nil, Val: 3}
	a.Next = b
	b.Next = c
	//d := reserv(a)

	for a != nil{
		fmt.Println(a.Val)
		a = a.Next
	}
}