# 100. Same Tree

**Difficulty:** Easy

Given the roots of two binary trees, determine whether they are structurally identical and have the same node values.

**LeetCode Go API:** `func isSameTree(p *TreeNode, q *TreeNode) bool`

## Examples

- `[1,2,3]`, `[1,2,3]` → `true`
- `[1,2]`, `[1,null,2]` → `false`
- `[1,2,1]`, `[1,1,2]` → `false`

## Constraints

- The number of nodes in both trees is in the range `[0, 100]`.
- `-10^4 <= Node.val <= 10^4`
