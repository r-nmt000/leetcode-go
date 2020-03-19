package main

func main() {

}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func isUnivalTree(root *TreeNode) bool {
	isLeftCorrect := (root.Left == nil ||
		(root.Val == root.Left.Val && isUnivalTree(root.Left)))

	isRightCorrect := (root.Right == nil ||
		(root.Val == root.Right.Val && isUnivalTree(root.Right)))

	return isLeftCorrect && isRightCorrect
}
