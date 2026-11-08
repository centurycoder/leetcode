package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) (total int) {
	sum := 0
	numbers := []int{}
	preorderTraverse(root, numbers, &sum)
	return sum
}

func preorderTraverse(root *TreeNode, numbers []int, sum *int) {
	if root == nil {
		return
	}
	numbers = append(numbers, root.Val)
	if root.Left == nil && root.Right == nil {
		pathSum := 0
		for _, digit := range numbers {
			pathSum = pathSum*10 + digit
		}
		*sum = *sum + pathSum
	}
	if root.Left != nil {
		preorderTraverse(root.Left, numbers, sum)
	}
	if root.Right != nil {
		preorderTraverse(root.Right, numbers, sum)
	}
}
