package main


func main() {
	depth2 := []*Node{{5, nil}, {6, nil}}
	depth1 := []*Node{{3, depth2}, {2, nil}, {4, nil}}
	top := Node{1, depth1}
	maxDepth(&top)
}

type Node struct {
	Val int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	maxDepth := 1
	getMaxDepth(root, 1, &maxDepth)
	return maxDepth
}

func getMaxDepth(node *Node, currentDepth int, maxDepth *int) {
	if currentDepth > *maxDepth {
		*maxDepth = currentDepth
	}
	for _, v := range node.Children {
		getMaxDepth(v, currentDepth+1, maxDepth)
	}
}
