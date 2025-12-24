package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	delNode := head
	currentNode := head

	for i := 0; i < n; i++ {
		if currentNode.Next == nil {
			return head.Next
		}

		currentNode = currentNode.Next
	}

	for currentNode.Next != nil {
		delNode = delNode.Next
		currentNode = currentNode.Next
	}

	delNode.Next = delNode.Next.Next

	return head
}

func main() {
	n := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	res := removeNthFromEnd(n, 3)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
