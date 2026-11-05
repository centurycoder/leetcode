package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]

	// find root position for in order
	rootPos := 0
	for i := range inorder {
		if inorder[i] == rootVal {
			rootPos = i
			break
		}
	}

	inorderLeft := inorder[:rootPos]
	preorderLeft := preorder[1 : 1+len(inorderLeft)]
	inorderRight := inorder[rootPos+1:]
	preorderRight := preorder[1+len(inorderLeft):]

	left := buildTree(preorderLeft, inorderLeft)
	right := buildTree(preorderRight, inorderRight)
	root := &TreeNode{
		Val:   preorder[0],
		Left:  left,
		Right: right,
	}
	return root
}
