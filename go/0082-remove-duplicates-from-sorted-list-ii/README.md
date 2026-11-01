# 82. Remove Duplicates from Sorted List II

**Difficulty:** Medium

Given the head of a sorted singly linked list, remove every value that occurs more than once. Return only nodes whose values are distinct in the original list.

**LeetCode Go API:** `func deleteDuplicates(head *ListNode) *ListNode`

## Examples

- `[1,2,3,3,4,4,5]` → `[1,2,5]`
- `[1,1,1,2,3]` → `[2,3]`

## Constraints

- The number of nodes is in the range `[0, 300]`.
- `-100 <= Node.val <= 100`
- The list is sorted in ascending order.

## Approach

Use a dummy node and a tail for the retained list. A value with an adjacent equal node begins a duplicate run; skip the entire run and connect the retained tail to the first different node. Otherwise, retain the current node. This handles duplicate runs at the head and tail without special cases.

```go
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := ListNode{Next: head}
	tail := &dummy

	for current := head; current != nil; {
		if current.Next != nil && current.Val == current.Next.Val {
			duplicate := current.Val
			for current != nil && current.Val == duplicate {
				current = current.Next
			}
			tail.Next = current
		} else {
			tail = current
			current = current.Next
		}
	}

	return dummy.Next
}
```

The traversal reuses `current` to skip a duplicate run, saving the run's value before advancing. This avoids a separate duplicate-detected flag and saved `next` pointer. `tail` advances only for a retained node; after a duplicate run, `tail.Next = current` reconnects the retained list to the next distinct value. The local value `dummy` provides a predecessor even when the head is removed.

- Time: O(n).
- Auxiliary space: O(1).

## Run tests

From the repository root:

```sh
go test ./go/0082-remove-duplicates-from-sorted-list-ii
```
