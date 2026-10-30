# 138. Copy List with Random Pointer

**Difficulty:** Medium

Deep-copy a linked list whose nodes have `Next` and `Random` pointers. Each random pointer is either `nil` or points to a node in the original list. The copied list must contain exactly one new node per original node, with corresponding values and pointer relationships; no copied pointer may reference an original node.

**LeetCode Go API:** `func copyRandomList(head *Node) *Node`

## Examples

- `[[7,null],[13,0],[11,4],[10,2],[1,0]]` → `[[7,null],[13,0],[11,4],[10,2],[1,0]]`
- `[[1,1],[2,1]]` → `[[1,1],[2,1]]`
- `[[3,null],[3,0],[3,null]]` → `[[3,null],[3,0],[3,null]]`

## Constraints

- `0 <= n <= 1000`
- `-10^4 <= Node.Val <= 10^4`
- `Node.Random` is `nil` or points to a node in the list.

## Analysis of the submitted implementation

The solution builds the copied `Next` chain in one pass while recording each original node’s address-to-copy mapping. A second pass uses that map to set each copied node’s `Random` pointer. Because all copies are recorded before random pointers are assigned, forward, backward, and self-references are handled. Looking up a `nil` pointer in the map returns `nil`, preserving null pointers. The algorithm takes expected O(n) time and O(n) auxiliary space.

The implementation is correct, but the first pass has unnecessary bookkeeping: `prevOld` is reassigned inside the loop but never read again, and `curNew = curNew.Next` reads the `Next` field of a newly allocated node before it is linked, so that assignment is always `nil`. Explicit `Next: nil` and `Random: nil` initializers are also redundant. Handling the head separately adds more state than needed.

## Simpler map-based implementation

Create all copied nodes first, then assign both pointers in a second pass. This avoids special-casing the head and makes the map’s role explicit. A missing map entry for a `nil` pointer evaluates to `nil`.

```go
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	clones := make(map[*Node]*Node)
	for original := head; original != nil; original = original.Next {
		clones[original] = &Node{Val: original.Val}
	}

	for original := head; original != nil; original = original.Next {
		clone := clones[original]
		clone.Next = clones[original.Next]
		clone.Random = clones[original.Random]
	}

	return clones[head]
}
```

## Run tests

From the repository root:

```sh
go test ./go/0138-copy-list-with-random-pointer
```
