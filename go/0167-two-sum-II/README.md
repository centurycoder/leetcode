# 167. Two Sum II — Input Array Is Sorted

**Difficulty:** Medium

Given a sorted integer array `numbers` and a target, return the 1-based indices of the two distinct elements whose sum is the target. The input guarantees exactly one solution.

**LeetCode Go signature:** `func twoSum(numbers []int, target int) []int`

## Example

- `numbers = [2, 7, 11, 15]`, `target = 9` → `[1, 2]`

## Approach: two pointers from opposite ends

Initialize `i` at the first element and `j` at the last. Compare `numbers[i] + numbers[j]` with `target`:

- If the sum is too small, increment `i`. Keeping this left value while moving `j` left would only make the sum smaller.
- If the sum is too large, decrement `j`. Keeping this right value while moving `i` right would only make the sum larger.
- If the sum matches, return `[i + 1, j + 1]` for the required 1-based indices.

Each step discards a pointer value that cannot be part of a solution on the remaining search interval. Both pointers only move inward, so the scan takes `O(n)` time and `O(1)` extra space.

## Comparison with the previous implementation

The previous version started `j` at `1`, scanned it forward for each `i`, then set `j = i + 1` when advancing `i`. The sorted order let it skip some pairs, but resetting `j` discarded its progress: many elements were revisited for each new `i`. In the worst case, this repeated scan takes `O(n^2)` time.

The key improvement is not just using two indexes; it is making both indexes move monotonically across the whole array. Starting at opposite ends lets each comparison rule out an entire row or column of candidate pairs without rescanning it.

## Lessons to carry forward

- Let input structure guide the algorithm: sorted order makes pointer movement informative.
- Prove each pointer move safe. A smaller sum rules out the current left value; a larger sum rules out the current right value.
- Count total index movement, not loop statements. One loop can still do quadratic work if an index is repeatedly reset.
- Preserve progress when scanning: monotonic pointers bound each element's participation and avoid revisiting candidates.

## Run tests

From the repository root:

```sh
go test ./go/0167-two-sum-II
```
