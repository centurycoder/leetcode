# 129. Sum Root to Leaf Numbers

**Difficulty:** Medium

Given the root of a binary tree containing digits from `0` to `9`, return the sum of all root-to-leaf numbers. Each path represents a number formed by concatenating its node values: `1 -> 2 -> 3` represents `123`. A leaf has no children.

**LeetCode Go API:** `func sumNumbers(root *TreeNode) int`

## Examples

- `[1,2,3]` → `25` (`12 + 13`)
- `[4,9,0,5,1]` → `1026` (`495 + 491 + 40`)

## My Approach

Use recursive depth-first traversal, carrying a slice of digits for the current root-to-node path. `sumNumbers` creates an empty path and a local total, then calls `preorderTraverse` with a pointer to that total.

- Append the current node's digit to the path.
- At a leaf, convert the path to a number by repeatedly applying `pathSum = pathSum*10 + digit`, then add it to the total.
- Recurse into the left and right children with the current path. A `nil` node contributes nothing.

`preorderTraverse` processes the current node before its left and right children, matching preorder DFS.

Go passes the slice header by value. Appending in a child does not change the parent's slice length, so explicit backtracking is unnecessary here. Sibling calls may reuse and overwrite the backing array beyond the parent's path prefix; this is safe because each leaf's number is consumed immediately rather than saving the slices.

The traversal visits every node once, but each leaf scans its whole path again to construct its number. For `n` nodes, `L` leaves, and height `h`, time is `O(n + Lh)` in the worst case. Auxiliary space is `O(h)` for the recursion stack and live path storage; appending can also allocate backing arrays.

## Standard Approach: DFS with a Running Number

Carry the number represented by the path so far instead of a slice of digits. At each node, compute `current = prefix*10 + node.Val`. At a leaf, return `current`; at an internal node, return the sum of the results from its children. A missing child returns `0`, so only complete root-to-leaf paths contribute.

```text
sumNumbers(root):
    return dfs(root, 0)

dfs(node, prefix):
    if node is nil:
        return 0
    current = prefix * 10 + node.value
    if node.left is nil and node.right is nil:
        return current
    return dfs(node.left, current) + dfs(node.right, current)
```

Passing `current` by value keeps each branch independent. Returning subtree sums removes the need for a shared total pointer.

Time: `O(n)`, since each node performs constant work. Auxiliary space: `O(h)` for the recursion stack. Compared with my approach, this avoids rescanning each leaf's path and allocating digit slices while retaining the same recursive traversal.

## Constraints

- The number of nodes is in the range `[1, 1000]`.
- `0 <= Node.val <= 9`
- The depth of the tree will not exceed `10`.
- The answer fits in a 32-bit integer.

## Local Tests

Run the implementation against the table-driven tests:

```sh
go test ./go/0129-sum-root-to-leaf-numbers
```

Tests cover both examples, single digits, leading and interior zeros, one-child paths, unequal path lengths, and a ten-digit number.
