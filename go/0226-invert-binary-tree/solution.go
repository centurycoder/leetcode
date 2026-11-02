package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) (result *TreeNode) {
	if root == nil {
		return nil
	}
	newLeft := invertTree(root.Left)
	newRight := invertTree(root.Right)
	root.Left = newRight
	root.Right = newLeft

	return root
}
