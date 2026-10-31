package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var cursor, firstLT, lastLT, firstGTE, lastGTE *ListNode
	for cursor = head; cursor != nil; {
		if cursor.Val < x {
			if firstLT == nil {
				firstLT = cursor
				lastLT = cursor
				cursor = cursor.Next
				lastLT.Next = nil
				continue
			} else {
				lastLT.Next = cursor
				lastLT = cursor
				cursor = cursor.Next
				lastLT.Next = nil
				continue
			}
		} else {
			if firstGTE == nil {
				firstGTE = cursor
				lastGTE = cursor
				cursor = cursor.Next
				lastGTE.Next = nil
				continue
			} else {
				lastGTE.Next = cursor
				lastGTE = cursor
				cursor = cursor.Next
				lastGTE.Next = nil
				continue
			}
		}
	}
	if lastLT != nil {
		lastLT.Next = firstGTE
	}
	if firstLT == nil {
		firstLT = firstGTE
	}
	return firstLT
}
