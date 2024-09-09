package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func maxTwinSum(head *ListNode) int {
	if head == nil || head.Next == nil {
		return 0
	}

	mid := findMiddle(head)

	secondHalf := reverseList(mid.Next)

	mid.Next = nil

	maxSum := 0
	firstHalf := head
	for secondHalf != nil {
		sum := firstHalf.Val + secondHalf.Val
		if sum > maxSum {
			maxSum = sum
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}

	return maxSum
}

func main() {
	head := &ListNode{Val: 5}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}

	result := maxTwinSum(head)
	fmt.Println("Maximum twin sum:", result)
}
