package main

import "fmt"

func main() {
	node7 := TreeNode{Val: 7, Left: nil, Right: nil}
	node8 := TreeNode{Val: 8, Left: nil, Right: nil}
	node5 := TreeNode{Val: 5, Left: nil, Right: nil}
	node4 := TreeNode{Val: 4, Left: &node7, Right: nil}
	node2 := TreeNode{Val: 2, Left: &node4, Right: &node5}
	node6 := TreeNode{Val: 6, Left: nil, Right: &node8}
	node3 := TreeNode{Val: 3, Left: nil, Right: &node6}
	node1 := TreeNode{Val: 1, Left: &node2, Right: &node3}
	fmt.Println(deepestLeavesSum(&node1))

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	hierarchySum := make(map[int]int)
	max := 0
	searchNode(root, 0, hierarchySum, &max)
	return hierarchySum[max]
}

func searchNode(node *TreeNode, hierarchy int, hierarchySum map[int]int, max *int) {
	if node == nil {
		return
	}
	if *max < hierarchy {
		*max = hierarchy
	}
	hierarchySum[hierarchy] = hierarchySum[hierarchy] + node.Val

	searchNode(node.Left, hierarchy+1, hierarchySum, max)
	searchNode(node.Right, hierarchy+1, hierarchySum, max)
}
