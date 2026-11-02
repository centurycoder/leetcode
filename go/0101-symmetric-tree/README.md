# 101. Symmetric Tree

**Difficulty:** Easy

Given the root of a binary tree, determine whether it is a mirror of itself around its center.

**LeetCode Go API:** `func isSymmetric(root *TreeNode) bool`

## Examples

- `[1,2,2,3,4,4,3]` → `true`
- `[1,2,2,null,3,null,3]` → `false`

## Approach

Check mirror pairs recursively. After handling a nil root and a single-node root, compare the root's left and right subtrees with `isMirror`. For nodes `a` and `b`, they mirror each other when both are nil, or when both exist, their values match, `a.Left` mirrors `b.Right`, and `a.Right` mirrors `b.Left`.

The traversal visits each node at most once, so time is `O(n)`. Recursion uses `O(h)` stack space, where `h` is the tree height (`O(n)` in the worst case).

## Key pitfalls

- Compare opposite sides: left-left against right-right, and left-right against right-left. Comparing matching child directions checks whether the subtrees are equal, not mirrored.
- Check both values and shape. Equal values do not compensate for children on the wrong side.
- Treat two missing nodes as a matching pair, but one missing node as a mismatch.
- A single-node tree is symmetric. A root with only one child is not.

## Constraints

- The number of nodes is in the range `[1, 1000]`.
- `-100 <= Node.val <= 100`
