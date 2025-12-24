// May be solved using multiple merging of 2 lists
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinVal struct {
	Idx int
	Val int
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}

	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	current := head

	for {
		minVal := &MinVal{
			Idx: -1,
			Val: 0,
		}

		// getting first non-nil value
		for i, list := range lists {
			if list != nil {
				minVal = &MinVal{
					Idx: i,
					Val: list.Val,
				}

				break
			}
		}

		// finding smallest value
		for i, list := range lists {
			if list != nil {
				if minVal.Idx != i && list.Val < minVal.Val {
					minVal.Val = list.Val
					minVal.Idx = i
				}
			}
		}

		if minVal.Idx == -1 {
			break
		}

		// appending node res
		current.Next = &ListNode{
			Val:  minVal.Val,
			Next: nil,
		}
		current = current.Next

		// popping existing node
		lists[minVal.Idx] = lists[minVal.Idx].Next
	}

	return head.Next
}

func main() {
	lN1 := &ListNode{
		Val:  1,
		Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}},
	}

	lN2 := &ListNode{
		Val:  1,
		Next: nil,
	}

	lN3 := &ListNode{
		Val:  1,
		Next: &ListNode{Val: 2, Next: &ListNode{Val: 10, Next: nil}},
	}

	res := mergeKLists([]*ListNode{lN1, lN2, lN3})

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
