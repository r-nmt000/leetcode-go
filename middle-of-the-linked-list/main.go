package main

import "fmt"

func main() {
	a := ListNode{3, nil}
	b := ListNode{2, &a}
	c := ListNode{1, &b}
	d := ListNode{ 1, nil}
	fmt.Println(middleNode(&d))
	fmt.Println(c)
}

type ListNode struct {
	Val int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	arry := make([]ListNode, 100)
	i := 0
	for ; head != nil; i++ {
		arry[i] = *head
		head = head.Next
	}
	return &arry[i/2]
}
