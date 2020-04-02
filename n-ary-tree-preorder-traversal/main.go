package main

func main() {
	a := Node{5, nil}
	b := Node{6, nil}
	c := Node{2, nil}
	d := Node{4, nil}
	e := Node{3, []*Node{&a, &b}}
	f := Node{1, []*Node{&e, &c, &d}}
	preorder(&f)

}

type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {
	result := []int{}
	preorderRecursive(root, result)
	return result
}

func preorderRecursive(node *Node, result []int) {
	if node.Children == nil {
		result = append(result, node.Val)
		return
	}
	for _,v := range node.Children {
		preorderRecursive(v, result)
	}
}

