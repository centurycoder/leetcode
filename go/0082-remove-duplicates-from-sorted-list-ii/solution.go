package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var dummy, prev, cur, next *ListNode
	dummy = &ListNode{}
	dummy.Next = head

	dupDetected := false

	prev = dummy
	for cur = head; cur != nil; {
		next = cur.Next
		for next != nil && cur.Val == next.Val {
			dupDetected = true
			cur = next
			next = next.Next
		}
		if dupDetected {
			prev.Next = next
			dupDetected = false
		} else {
			prev = cur
		}
		cur = next
	}
	return dummy.Next
}
