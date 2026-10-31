# 86. Partition List

**Difficulty:** Medium

Given the head of a singly linked list and a value `x`, partition the list so that nodes with values less than `x` come before nodes with values greater than or equal to `x`. Preserve the original relative order within both partitions.

**LeetCode Go API:** `func partition(head *ListNode, x int) *ListNode`

## Examples

- `[1,4,3,2,5,2]`, `x = 3` → `[1,2,2,4,3,5]`
- `[2,1]`, `x = 2` → `[1,2]`

## Constraints

- The number of nodes is in the range `[0, 200]`.
- `-100 <= Node.val <= 100`
- `-200 <= x <= 200`

## Suggested solution (reference)

The implementation in `solution.go` is kept unchanged. This alternative uses two local dummy head nodes so appending the first node and later nodes follows the same path.

```go
func partition(head *ListNode, x int) *ListNode {
	var lessDummy, greaterOrEqualDummy ListNode
	lessTail := &lessDummy
	greaterOrEqualTail := &greaterOrEqualDummy

	for cursor := head; cursor != nil; {
		// Save the remaining list before detaching the current node.
		next := cursor.Next
		cursor.Next = nil

		if cursor.Val < x {
			lessTail.Next = cursor
			lessTail = cursor
		} else {
			greaterOrEqualTail.Next = cursor
			greaterOrEqualTail = cursor
		}

		cursor = next
	}

	lessTail.Next = greaterOrEqualDummy.Next
	return lessDummy.Next
}
```

- Appending to each tail preserves relative order and reuses the original nodes.
- Saving `next` before detaching keeps traversal separate from link updates.
- Dummy heads remove empty-partition special cases, including when the input is empty.
- Time: O(n). Auxiliary space: O(1).

## Run tests

From the repository root:

```sh
go test ./go/0086-partition-list
```
