package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	var prev, cur, tailLeft, tailMiddle *ListNode
	pos := 0
	dummyNode := ListNode{}

	for cur = head; cur != nil; {
		pos++
		next := cur.Next

		if pos == left {
			tailLeft = prev
			tailMiddle = cur
			cur.Next = nil
			if tailLeft == nil {
				tailLeft = &dummyNode
				tailLeft.Next = cur
			}
		} else if pos > left && pos <= right {
			if tailLeft != nil {
				cur.Next = tailLeft.Next
				tailLeft.Next = cur
			}

		} else if pos == right+1 {
			tailMiddle.Next = cur
		}
		prev = cur
		cur = next
	}

	if tailLeft == &dummyNode {
		return tailLeft.Next
	}
	return head
}
