package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	startNode := &ListNode{}
	curr := startNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return startNode.Next
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			}},
	}

	l2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	res := mergeTwoLists(l1, l2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
