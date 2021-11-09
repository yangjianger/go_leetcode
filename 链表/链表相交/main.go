package main

import "fmt"

type Node struct {
	val int
	next *Node
}

func xjLinked(a *Node, b *Node) *Node{

	if a ==  nil || b == nil{
		return nil
	}

	curA := a
	curB := b

	for curA != curB{
		if curA == nil{
			curA = b
		}else{
			curA = curA.next
		}

		if curB == nil{
			curB = a
		}else{
			curB = curB.next
		}
	}

	return curA
}

func main(){
	a := &Node{1, nil}
	b := &Node{2, nil}
	c := &Node{3, nil}
	a.next = b
	b.next = c

	d := &Node{4, nil}
	d.next = b
	b.next = c

	newHead := xjLinked(a, d)

	cur := newHead

	for cur != nil{
		fmt.Println(cur.val)
		cur = cur.next
	}
}
