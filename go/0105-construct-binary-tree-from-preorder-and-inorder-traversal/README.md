# 105. Construct Binary Tree from Preorder and Inorder Traversal

**Difficulty:** Medium

Given the preorder and inorder traversals of a binary tree with unique values, construct and return the tree.

**LeetCode Go API:** `func buildTree(preorder []int, inorder []int) *TreeNode`

## Examples

- `preorder = [3,9,20,15,7]`, `inorder = [9,3,15,20,7]` → `[3,9,20,null,null,15,7]`
- `preorder = [-1]`, `inorder = [-1]` → `[-1]`

## Approach

Preorder always places the current subtree's root first, but does not reveal the sizes of its left and right subtrees. Find that root in inorder: values before it belong to the left subtree, and values after it belong to the right. The root's inorder position gives the left-subtree size, which lets us split the remaining preorder values into the matching left and right subtrees. Repeat recursively. The standard implementation below maps inorder values to indices, giving `O(n)` time and `O(n)` auxiliary space.

## Standard Go Implementation

```go
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderIndex := make(map[int]int, len(inorder))
	for i, val := range inorder {
		inorderIndex[val] = i
	}

	preorderIndex := 0
	return buildTreeRange(preorder, 0, len(inorder), &preorderIndex, inorderIndex)
}

func buildTreeRange(preorder []int, start, end int, preorderIndex *int, inorderIndex map[int]int) *TreeNode {
	if start >= end {
		return nil
	}

	val := preorder[*preorderIndex]
	*preorderIndex = *preorderIndex + 1
	root := &TreeNode{Val: val}
	pivot := inorderIndex[val]
	root.Left = buildTreeRange(preorder, start, pivot, preorderIndex, inorderIndex)
	root.Right = buildTreeRange(preorder, pivot+1, end, preorderIndex, inorderIndex)
	return root
}
```

This uses `O(n)` time and `O(n)` auxiliary space. `TreeNode` is provided by LeetCode.

## Constraints

- `1 <= preorder.length <= 3000`
- `inorder.length == preorder.length`
- `-3000 <= preorder[i], inorder[i] <= 3000`
- Values are unique, and both traversals describe the same tree.
