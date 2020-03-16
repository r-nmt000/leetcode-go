package main

import "fmt"

func main() {
	node1 := TreeNode{Val: 1, Left: nil, Right: nil}
	node2 := TreeNode{Val: 2, Left: &node1, Right: nil}
	node4 := TreeNode{Val: 4, Left: nil, Right: nil}
	node3 := TreeNode{Val: 3, Left: &node2, Right: &node4}
	node7 := TreeNode{Val: 7, Left: nil, Right: nil}
	node9 := TreeNode{Val: 9, Left: nil, Right: nil}
	node8 := TreeNode{Val: 8, Left: &node7, Right: &node9}
	node6 := TreeNode{Val: 6, Left: nil, Right: &node8}
	root := TreeNode{Val: 5, Left: &node3, Right: &node6}
	fmt.Println(increasingBST(&root))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	vals := []int{}
	increasingBSTRecursive(root, &vals)
	ans := &TreeNode{Val:0}
	cur := ans
	for _, v := range vals {
		cur.Right = &TreeNode{Val:v}
		cur = cur.Right
	}
	return ans.Right
}

func increasingBSTRecursive(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	increasingBSTRecursive(root.Left, vals)
	*vals = append(*vals, root.Val)
	increasingBSTRecursive(root.Right, vals)
}
