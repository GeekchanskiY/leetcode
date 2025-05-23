package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// https://live-scripts.blogspot.com/2013/01/blog-post_16.html - it would be fun to implement the second variant

func main() {
	listNode1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	listNode2 := &ListNode{
		Val:  2,
		Next: listNode1,
	}
	listNode3 := &ListNode{
		Val:  3,
		Next: listNode2,
	}
	listNode4 := &ListNode{
		Val:  4,
		Next: listNode3,
	}
	listNode5 := &ListNode{
		Val:  5,
		Next: listNode4,
	}
	listNode6 := &ListNode{
		Val:  6,
		Next: listNode5,
	}
	listNode7 := &ListNode{
		Val:  7,
		Next: listNode6,
	}
	listNode8 := &ListNode{
		Val:  8,
		Next: listNode7,
	}
	listNode9 := &ListNode{
		Val:  9,
		Next: listNode8,
	}
	listNode10 := &ListNode{
		Val:  10,
		Next: listNode9,
	}
	listNode1.Next = listNode6

	fmt.Println(hasCycle(listNode10))

	listNodeSelfRef := &ListNode{
		Val: 1,
	}
	listNodeSelfRef.Next = listNodeSelfRef
	fmt.Println(hasCycle(listNodeSelfRef))

}
