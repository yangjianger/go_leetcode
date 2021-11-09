package main

import "fmt"

type Node struct {
	val int
	next *Node
}

func addLinked(a *Node, b *Node) *Node{
	if a == nil {
		return b
	}

	if b == nil{
		return a
	}

	var (
		dummyHead = &Node{0, nil}
		tail = dummyHead
		cur = 0
		v1 = 0
		v2 = 0
		sum = 0
	)

	for a != nil ||  b != nil{
		v1 = 0
		v2 = 0
		if a != nil{
			v1 = a.val
			a = a.next
		}
		if b != nil{
			v2 = b.val
			b = b.next
		}

		sum = v1 + v2 + cur
		cur = sum / 10

		tail.next = &Node{sum % 10 , nil}
		tail = tail.next
	}

	if cur > 0{
		tail.next = &Node{cur , nil}
	}

	return dummyHead.next

}

func main(){
	a := &Node{1, nil}
	b := &Node{2, nil}
	c := &Node{3, nil}
	a.next = b
	b.next = c

	a2 := &Node{7, nil}
	b2 := &Node{8, nil}
	c2 := &Node{9, nil}
	a2.next = b2
	b2.next = c2

	newHead  := addLinked(a, a2)
	cur := newHead

	for cur != nil{
		fmt.Println(cur.val)
		cur = cur.next
	}


}
