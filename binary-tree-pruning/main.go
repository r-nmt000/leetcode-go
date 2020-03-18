package main

import "fmt"

func main() {
	node1 := TreeNode{Val: 0, Left: nil, Right: nil}
	node2 := TreeNode{Val: 1, Left: nil, Right: nil}
	node3 := TreeNode{Val: 0, Left: &node1, Right: &node2}
	node4 := TreeNode{Val: 1, Left: nil, Right: &node3}
	pruneTree(&node4)
	fmt.Println(node4)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if containsOne(root) {
		return root
	}
	return nil
}

func containsOne(node *TreeNode) bool {
	if node == nil { return false }
	isLeftValid := containsOne(node.Left)
	isRightValid := containsOne(node.Right)
	if !isLeftValid { node.Left = nil }
	if !isRightValid { node.Right = nil }
	return node.Val == 1 || isLeftValid || isRightValid
}
