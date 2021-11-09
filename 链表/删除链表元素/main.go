package main

import "fmt"

type Node struct {
	val int
	next *Node
}

func removeNode(root *Node, val int) *Node{

	dummyHead := &Node{val: 0, next: nil}
	tail := dummyHead

	for root != nil {
		if root.val != val{
			tail.next = root
			tail = tail.next
		}

		root = root.next
	}

	tail.next = nil

	return dummyHead.next
}

func main(){
	a := &Node{1, nil}
	b := &Node{2, nil}
	c := &Node{3, nil}
	d := &Node{4, nil}
	a.next = b
	b.next = c
	c.next = d

	root := removeNode(a, 2)
	cur := root
	for cur != nil{
		fmt.Println(cur.val)
		cur = cur.next
	}

}
