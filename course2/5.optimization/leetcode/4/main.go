package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	var sum int
	return convert(root, &sum)
}

func convert(node *TreeNode, sum *int) *TreeNode {
	if node == nil {
		return nil
	}

	convert(node.Right, sum)

	*sum += node.Val
	node.Val = *sum

	convert(node.Left, sum)

	return node
}

func inorderTraversal(root *TreeNode) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Print(root.Val, " ")
		inorderTraversal(root.Right)
	}
}

func main() {
	root := &TreeNode{Val: 4,
		Left: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 0},
			Right: &TreeNode{Val: 2,
				Right: &TreeNode{Val: 3}}},
		Right: &TreeNode{Val: 6,
			Left: &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7,
				Right: &TreeNode{Val: 8}}}}

	convertBST(root)

	inorderTraversal(root)
	fmt.Println()
}
