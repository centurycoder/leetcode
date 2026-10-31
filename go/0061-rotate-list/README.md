# 61. Rotate List

**Difficulty:** Medium

Given the head of a singly linked list, rotate the list to the right by `k` places.

**LeetCode Go API:** `func rotateRight(head *ListNode, k int) *ListNode`

## Examples

- `[1,2,3,4,5]`, `k = 2` → `[4,5,1,2,3]`
- `[0,1,2]`, `k = 4` → `[2,0,1]`

## Constraints

- The number of nodes is in the range `[0, 500]`.
- `-100 <= Node.val <= 100`
- `0 <= k <= 2 * 10^9`

## O(1)-space implementation (`rotateRight`)

`rotateRight` counts the nodes while finding the tail. It links the tail to the head to form a temporary cycle, then advances from the old tail by `n - (k % n)` links. The current node is the new tail; its next node is the new head. The function breaks the cycle before returning. When `k` is a multiple of the list length, it advances a full cycle and restores the original list.

```go
func rotateRight(head *ListNode, k int) *ListNode {
	var p *ListNode
	length := 0
	if head == nil {
		return head
	}
	for p = head; p.Next != nil; p = p.Next {
		length++
	}
	length++

	p.Next = head
	for shift := length - k%length; shift > 0; shift-- {
		p = p.Next
	}
	newHead := p.Next
	p.Next = nil
	return newHead
}
```

This implementation runs in **O(n)** time and uses **O(1)** extra space. The nil-list check precedes the modulo operation.

## Suggested O(1)-space solution

This is the alternative version suggested earlier. It reduces `k` before forming a cycle, returns immediately when no rotation is needed, and finds the new tail by advancing `n-k-1` links from the head. This code is an alternative to the `rotateRight` implementation above; use one version under the LeetCode API name, not both in the same Go file.

```go
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	tail := head
	length := 1
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	k %= length
	if k == 0 {
		return head
	}

	tail.Next = head
	newTail := head
	for i := 0; i < length-k-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}
```

It also runs in **O(n)** time and uses **O(1)** extra space. The early return avoids creating a temporary cycle when `k` is a multiple of the list length.

## Array-based alternative (`rotateRightWithArray`)

The array-based version stores every node pointer in a slice. After computing `k % n`, it uses index `n-k` as the new head, links the old tail to the old head, and cuts the list immediately before the new head. It also runs in **O(n)** time, but the slice uses **O(n)** extra space. This approach was inspired by the pointer-tracking idea in [`go/0138-copy-list-with-random-pointer`](../0138-copy-list-with-random-pointer/README.md): retain node pointers while traversing, then use them later. Here, the slice also preserves list order so indexes identify the rotation point.

## Run tests

From the repository root:

```sh
go test ./go/0061-rotate-list
```
