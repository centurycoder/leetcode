# 155. Min Stack

**Difficulty:** Medium

Design a stack that supports pushing, popping, reading the top value, and retrieving the minimum value in constant time.

**LeetCode Go API:** `Constructor() MinStack`, `Push(value int)`, `Pop()`, `Top() int`, `GetMin() int`

## Main idea

The current minimum is always stored at the stack top. To maintain that, each entry stores both its value and the minimum seen at that stack depth. `Push` derives the entry's minimum from the incoming value and the previous top; `Pop` reveals the previous entry and its minimum. Every operation takes O(1) time.

## Example

After pushing `-2`, `0`, and `-3`, `GetMin()` returns `-3`. After popping, `Top()` returns `0` and `GetMin()` returns `-2`.

## Constraints

- Values are in the signed 32-bit integer range.
- `Pop`, `Top`, and `GetMin` are called only when the stack is non-empty.
- At most `3 * 10^4` operations are made.

## Run tests

From the repository root:

```sh
go test ./go/0155-min-stack
```
