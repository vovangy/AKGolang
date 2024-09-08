package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left, result)
	*result = append(*result, root.Val)
	inorderTraversal(root.Right, result)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

func balanceBST(root *TreeNode) *TreeNode {
	var nums []int
	inorderTraversal(root, &nums)
	return sortedArrayToBST(nums)
}

func printInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	printInOrder(root.Left)
	fmt.Print(root.Val, " ")
	printInOrder(root.Right)
}

func main() {
	root := &TreeNode{Val: 4,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 6,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7}},
	}

	balancedRoot := balanceBST(root)
	fmt.Print("Balanced BST: ")
	printInOrder(balancedRoot)
}
