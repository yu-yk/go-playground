package main

import (
	"fmt"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	nodes := []*ListNode{}
	curr := head
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	for i := len(nodes) - 1; i > (len(nodes)-1)/2; i-- {
		next := curr.Next
		curr.Next = nodes[i]
		curr.Next.Next = next
		curr = next
	}
	curr.Next = nil
}

func reverseListRecursion(head *ListNode) *ListNode {
	// 1 > 2 > 3 > nil
	if head == nil {
		return head
	} // edge case []
	if head.Next == nil {
		return head
	} // base case 1-> nil

	newHead := reverseListRecursion(head.Next)
	// newHead = 3 always tail of list
	// head = 2

	head.Next.Next = head
	// 1 > 2 > < 3

	// set itself point to nil
	head.Next = nil
	// 1 > 2 > nil
	//     ^
	//     3

	return newHead

	// next round:
	// head = 1
	// head.Next.Next = head
	// 1 > < 2
	//       ^
	//       3
	// head.Next = nil
	// nil < 1 < 2 < 3
}

func reverseListLoop(head *ListNode) *ListNode {
	if head == nil {
		return head
	} // edge case []
	if head.Next == nil {
		return head
	} // base case 1-> nil

	var prev *ListNode
	curr := head
	var next *ListNode
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// O(n) Cycle Detection
func findDuplicate(nums []int) int {
	fast, slow := nums[0], nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	p1, p2 := nums[0], fast
	for p1 != p2 {
		p1 = nums[p1]
		p2 = nums[p2]
	}
	return p1
}

func main() {
	// node5 := &ListNode{5, nil}
	node4 := &ListNode{4, nil}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	reorderList(node1)

	curr := node1
	i := 0
	for curr != nil || i > 10 {
		fmt.Print(curr.Val, " ")
		curr = curr.Next
		i++
	}
}
