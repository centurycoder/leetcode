# 112. Path Sum

**Difficulty:** Easy

Given the root of a binary tree and an integer `targetSum`, determine whether the tree has a root-to-leaf path whose node values sum to `targetSum`. A leaf has no children.

**LeetCode Go API:** `func hasPathSum(root *TreeNode, targetSum int) bool`

## Examples

- `[5,4,8,11,null,13,4,7,2,null,null,null,1]`, `22` → `true`
- `[1,2,3]`, `5` → `false`
- `[]`, `0` → `false`

## Suggested Solution

Use depth-first search while carrying the remaining sum. Return `false` for an empty node. At a leaf, return whether its value equals the remaining sum; otherwise subtract the current node's value and search either child. Checking only at leaves prevents a matching root-to-internal-node prefix from counting as a path.

```text
hasPathSum(node, remaining):
    if node is empty: return false
    if node is a leaf: return node.value == remaining
    next = remaining - node.value
    return hasPathSum(node.left, next) OR hasPathSum(node.right, next)
```

Time: `O(n)` for `n` nodes. Auxiliary space: `O(h)` recursion stack for tree height `h`.

## Constraints

- The number of nodes in the tree is in the range `[0, 5000]`.
- `-1000 <= Node.val <= 1000`
- `-1000 <= targetSum <= 1000`
