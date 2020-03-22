package main

import (
	"fmt"
	"strconv"
)

func main() {

	node1 := ListNode{1, nil}
	node2 := ListNode{ 0, &node1}
	node3 := ListNode{ 1, &node2}
	fmt.Println(getDecimalValue(&node3))

}

type ListNode struct {
	Val int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	str := getStrValue(head)
	result, _ := strconv.ParseInt(str, 2, 32)

	return int(result)
}

func getStrValue(node *ListNode) string {
	if node == nil { return ""}
	valStr := strconv.Itoa(node.Val)
	return valStr + getStrValue(node.Next)
}
