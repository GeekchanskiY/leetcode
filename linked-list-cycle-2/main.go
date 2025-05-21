package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}

	slow, fast := head, head.Next
	for {
		if fast.Next == nil {
			return nil
		}
		if fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next

		if fast.Next == slow {
			return slow
		}
	}

	return nil
}
