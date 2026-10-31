package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRightWithArray(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var nodes []*ListNode
	for p := head; p != nil; p = p.Next {
		nodes = append(nodes, p)
	}

	length := len(nodes)
	k = k % length
	if k == 0 {
		return head
	}

	offset := length - k
	nodes[length-1].Next = nodes[0]
	nodes[offset-1].Next = nil

	return nodes[offset]
}

func rotateRight(head *ListNode, k int) *ListNode {
	var p *ListNode
	length := 0
	if head == nil {
		return head
	}
	for p = head; p.Next != nil; p = p.Next {
		length++
	}
	// we quit early, so increase length by 1 to count the last node
	length++

	// temporarily cyclic the list
	p.Next = head

	for shift := length - k%length; shift > 0; shift-- {
		p = p.Next
	}
	newHead := p.Next
	p.Next = nil

	return newHead
}
