package solution

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) (connectedRoot *Node) {
	if root == nil {
		return nil
	}
	parentNodes := []*Node{root}
	for len(parentNodes) != 0 {
		childNodes := []*Node{}
		var prevNode *Node
		for _, parentNode := range parentNodes {
			if parentNode.Left != nil {
				childNodes = append(childNodes, parentNode.Left)
			}
			if parentNode.Right != nil {
				childNodes = append(childNodes, parentNode.Right)
			}
			if prevNode != nil {
				prevNode.Next = parentNode
			}
			prevNode = parentNode
		}
		parentNodes = childNodes
	}
	return root
}
