package main

/*
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。
*/

/*
解题思路：
	1.这里可以创建连个虚拟链表，分别为lHead和rHead，将小于x的放在lHead之后，大于等于x的放在rHead，最后 lTail.next=rHead 将两个链表组合
	2.这里有个注意点：在链表遍历完成之后，需要将rTail.next=nil,因为 可能出现  5->2->2->2   这种情况
 */

 type ListNode struct {
	Val int
	 Next *ListNode
 }
func partition(head *ListNode, x int) *ListNode {
	//这里如果不创建虚拟头结点，每次都会判断节点是否为空
	var (
		lHead = &ListNode{Val: -1, Next: nil}
		lTail = lHead
		rHead = &ListNode{Val: -1, Next: nil}
		rTail = rHead
		cur = head
	)

	for cur != nil {
		if cur.Val < x {
			lTail.Next = cur
			lTail = lTail.Next
		} else {
			rTail.Next = cur
			cur = cur.Next
		}
	}
	//在链表遍历完成之后，需要将rTail.next=nil,因为 可能出现  5->2->2->2   这种情况
	rTail.Next = nil
	lTail.Next = rHead.Next
	return lHead.Next
}