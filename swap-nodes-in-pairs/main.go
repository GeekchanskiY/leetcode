// https://leetcode.com/problems/swap-nodes-in-pairs

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first, second := head, head.Next
	first.Next, second.Next = swapPairs(head.Next.Next), first
	return second
}

func main() {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	newH := swapPairs(h)
	for newH != nil {
		fmt.Println(newH.Val)
		newH = newH.Next
	}
}
