package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMaxIndex(nums []int) int {
	maxIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIndex := findMaxIndex(nums)

	root := &TreeNode{
		Val: nums[maxIndex],
	}

	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])

	return root
}

func printTreeByLevel(root *TreeNode) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(*TreeNode)
		queue.Remove(element)

		fmt.Print(node.Val, " ")

		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	fmt.Println()
}

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}

	root := constructMaximumBinaryTree(nums)

	fmt.Print("Tree values by level: ")
	printTreeByLevel(root)
}
