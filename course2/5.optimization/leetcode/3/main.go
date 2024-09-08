package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	sum := 0

	for head != nil {
		if head.Val == 0 {
			if sum != 0 {
				current.Next = &ListNode{Val: sum}
				current = current.Next
				sum = 0
			}
		} else {
			sum += head.Val
		}
		head = head.Next
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	head := &ListNode{Val: 0,
		Next: &ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 0,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 4,
							Next: &ListNode{Val: 5,
								Next: &ListNode{Val: 0}}}}}}}}

	result := mergeNodes(head)
	printList(result)
}
