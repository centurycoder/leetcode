package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) (symmetric bool) {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && root.Right != nil {
		return isMirror(root.Left, root.Right)
	}

	return false
}

func isMirror(T, S *TreeNode) bool {
	if T == nil && S == nil {
		return true
	}

	if T != nil && S != nil {
		if isMirror(T.Left, S.Right) && isMirror(T.Right, S.Left) && T.Val == S.Val {
			return true
		}
	}

	return false
}
