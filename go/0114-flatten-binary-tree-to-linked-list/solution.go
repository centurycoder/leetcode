package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var prev *TreeNode

func flatten(root *TreeNode) {
	prev = nil
	_flatten(root)
}

func _flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Right != nil {
		_flatten(root.Right)
	}
	if root.Left != nil {
		_flatten(root.Left)
	}
	root.Right = prev
	root.Left = nil
	prev = root
}
