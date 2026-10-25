# 283. Move Zeroes

**Difficulty:** Easy

Given an integer slice `nums`, move every zero to the end while preserving the relative order of non-zero values. Modify the slice in place; do not make a copy.

**LeetCode Go signature:** `func moveZeroes(nums []int)`

## Examples

- `[0, 1, 0, 3, 12]` → `[1, 3, 12, 0, 0]`
- `[0]` → `[0]`

## Constraints

- `1 <= len(nums) <= 10^4`
- `-2^31 <= nums[i] <= 2^31 - 1`

## Implementation

The function compacts non-zero values in order, then fills the remaining positions with zeros. It runs in **O(n)** time and uses **O(1)** extra space.

## Run tests

From the repository root:

```sh
go test ./go/0283-move-zeroes
```

To run this problem's tests with verbose case names:

```sh
go test -v ./go/0283-move-zeroes
```

To run every Go problem's tests:

```sh
go test ./...
```

`solution_test.go` covers the examples, all-zero and no-zero inputs, and zeros at the beginning, between values, and at the end.
