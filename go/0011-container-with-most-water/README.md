# 11. Container With Most Water

**Difficulty:** Medium

Given an integer slice `height`, choose two lines that form a container with the x-axis. Return the maximum amount of water it can hold; the container cannot be slanted.

**LeetCode Go signature:** `func maxArea(height []int) int`

## Examples

- `[1, 8, 6, 2, 5, 4, 8, 3, 7]` → `49`
- `[1, 1]` → `1`

## Constraints

- `2 <= len(height) <= 10^5`
- `0 <= height[i] <= 10^4`

## Run tests

From the repository root:

```sh
go test ./go/0011-container-with-most-water
```
