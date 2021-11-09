package main

import "fmt"

type Node struct {
	val int
	next *Node
}

func reserveNode(root *Node) *Node{
	if root.next == nil{
		return root
	}

	newHead := reserveNode(root.next)
	root.next.next = root
	root.next = nil

	return newHead
}

func main(){
	a := &Node{1, nil}
	b := &Node{2, nil}
	c := &Node{3, nil}
	d := &Node{4, nil}
	a.next = b
	b.next = c
	c.next = d

	newHead := reserveNode(a)
	cur := newHead

	for cur != nil{
		fmt.Println(cur.val)
		cur = cur.next
	}

}
