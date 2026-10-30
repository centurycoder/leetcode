package solution

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var prevOld, curOld, prevNew, curNew, headNew *Node

	addrMapping := map[*Node]*Node{}

	if head == nil {
		return nil
	}

	prevOld = head
	curOld = head.Next

	headNew = &Node{
		Val:    prevOld.Val,
		Next:   nil,
		Random: nil,
	}
	prevNew = headNew
	addrMapping[prevOld] = prevNew

	for curOld != nil {
		curNew = &Node{
			Val:    curOld.Val,
			Next:   nil,
			Random: nil,
		}
		addrMapping[curOld] = curNew

		// move ahead for the old and new chain
		prevNew.Next = curNew
		prevNew = curNew
		curNew = curNew.Next
		prevOld = curOld
		curOld = curOld.Next
	}

	// loop again
	curOld = head
	curNew = headNew
	for curOld != nil && curNew != nil {
		curNew.Random = addrMapping[curOld.Random]
		curOld = curOld.Next
		curNew = curNew.Next
	}

	return headNew
}
