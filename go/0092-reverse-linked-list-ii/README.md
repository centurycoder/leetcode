# 92. Reverse Linked List II

**Difficulty:** Medium

Given the head of a singly linked list and integers `left` and `right`, reverse the nodes from position `left` through position `right` and return the list.

**LeetCode Go API:** `func reverseBetween(head *ListNode, left, right int) *ListNode`

## Examples

- `[1,2,3,4,5]`, `left = 2`, `right = 4` → `[1,4,3,2,5]`
- `[5]`, `left = 1`, `right = 1` → `[5]`

## Constraints

- `1 <= n <= 500`
- `-500 <= Node.val <= 500`
- `1 <= left <= right <= n`

## Standard iterative solution

Use a dummy node so the predecessor of the reversed range exists even when `left == 1`. Keep `before` at that predecessor and repeatedly move the node after `current` to the front of the range. `current` stays at the range's tail as nodes are moved, so the suffix remains connected without a separate reconnect step.

```go
func reverseBetween(head *ListNode, left, right int) *ListNode {
	dummy := ListNode{Next: head}
	before := &dummy
	for position := 1; position < left; position++ {
		before = before.Next
	}

	current := before.Next
	for i := 0; i < right-left; i++ {
		next := current.Next
		current.Next = next.Next
		next.Next = before.Next
		before.Next = next
	}

	return dummy.Next
}
```

- Time: O(n); auxiliary space: O(1).
- The loop runs `right - left` times, moving each selected node in front of the range.

## Run tests

From the repository root:

```sh
go test ./go/0092-reverse-linked-list-ii
```
