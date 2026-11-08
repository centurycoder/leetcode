# 114. Flatten Binary Tree to Linked List

**Difficulty:** Medium

Given the root of a binary tree, flatten it into a linked list in-place. The list uses the same `TreeNode` type: each node's `Right` pointer points to the next node, and each `Left` pointer is `nil`. The order must match a preorder traversal.

**LeetCode Go API:** `func flatten(root *TreeNode)`

## Examples

- `[1,2,5,3,4,null,6]` → `[1,null,2,null,3,null,4,null,5,null,6]`
- `[]` → `[]`
- `[0]` → `[0]`

## My Approach

Traverse in reverse preorder: right subtree, then left subtree, then the current node. Keep `prev` as the head of the already-flattened suffix. After visiting both children, set `root.Right = prev`, clear `root.Left`, and update `prev = root`. This prepends each node to the suffix, so the resulting right-only chain is in preorder. Resetting `prev` at the public entry point makes separate calls independent.

Time: `O(n)` for `n` nodes. Auxiliary space: `O(h)` recursion stack for tree height `h`.

## Suggested Solution (in-place, `O(1)` extra space)

Walk the tree along the resulting right chain. When a node has a left subtree, find that subtree's rightmost node, append the node's original right subtree there, then move the left subtree to the right and clear the left pointer. This preserves preorder: left subtree first, followed by the original right subtree.

```text
flatten(root):
    current = root
    while current exists:
        if current.left exists:
            predecessor = current.left
            while predecessor.right exists:
                predecessor = predecessor.right
            predecessor.right = current.right
            current.right = current.left
            current.left = nil
        current = current.right
```

### Why the rewiring works

Preorder visits the current node, then every node in its left subtree, then every node in its right subtree. So when `current.left` exists, move that left subtree to the right, but first preserve the original right subtree by attaching it after the left subtree's rightmost node:

```text
predecessor.right = current.right  // keep the original right subtree
current.right = current.left       // visit the left subtree next
current.left = nil
```

The predecessor is found by following `Right` pointers from `current.left` until reaching a node whose `Right` is `nil`. That empty link is where the original right subtree belongs. The order matters: attaching the old right subtree first prevents losing it when `current.right` is overwritten.

For example, start with:

```text
        1
       / \
      2   5
     / \   \
    3   4   6
```

At node `1`, following right pointers from its left child `2` reaches `4`. Set `4.right = 5`, then move `2` to `1.right` and clear `1.left`. Next, at node `2`, set `3.right = 4`, move `3` to `2.right`, and clear `2.left`. The resulting chain is `1 → 2 → 3 → 4 → 5 → 6`.

The loop then advances along `current.right`, processing nodes in preorder. Nodes without a left child need no rewiring. The predecessor search follows right links only; it does not mean the visually farthest-right node in the diagram.

Each node is processed in preorder, and the predecessor links splice the original right subtree after the left subtree without allocating a stack or collection. Time: `O(n)`. Auxiliary space: `O(1)`.

## Constraints

- The number of nodes in the tree is in the range `[0, 2000]`.
- `-100 <= Node.val <= 100`

**Follow-up:** Can you flatten the tree in-place with `O(1)` extra space?
