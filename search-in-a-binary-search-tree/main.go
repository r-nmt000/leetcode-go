package main

import "fmt"

type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	left := TreeNode{val: 3, Left: nil, Right: nil}
	right := TreeNode{val: 4, Left: nil, Right: nil}
	root := TreeNode{val: 1, Left: &left, Right: &right}
	fmt.Println(searchBST(&root, 5))
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if val == root.val {
		return root
	}
	if root.Left != nil {
		leftTree := searchBST(root.Left, val)
		if leftTree != nil {
			return leftTree
		}
	}
	if root.Right != nil {
		return searchBST(root.Right, val)
	}
	return nil
}
