# 15. 3Sum

**Difficulty:** Medium

Given an integer slice `nums`, return all distinct triplets whose values sum to zero. The order of the triplets and the values within each triplet do not matter.

**LeetCode Go signature:** `func threeSum(nums []int) [][]int`

## Examples

- `[-1, 0, 1, 2, -1, -4]` → `[[-1, -1, 2], [-1, 0, 1]]`
- `[0, 1, 1]` → `[]`
- `[0, 0, 0]` → `[[0, 0, 0]]`

## Constraints

- `3 <= len(nums) <= 3000`
- `-10^5 <= nums[i] <= 10^5`

## Approach and debugging notes

### Sorting and two pointers

Sorting is not the only possible approach, but it makes the standard two-pointer solution efficient and straightforward. For each fixed `i`, the remaining values are ordered: if the sum is too small, move `j` right; if it is too large, move `k` left. Equal values are adjacent, which makes duplicate handling possible. The problem returns values rather than original indices, so sorting does not lose information the answer needs.

### Why deduplication happens in two places

- **At `i`:** if `nums[i] == nums[i-1]`, skip this anchor. The earlier equal value already considered every triplet that can start with this value; using this anchor again would repeat those triplets.
- **After finding a zero sum:** append the triplet, move both `j` and `k` inward, then skip adjacent values equal to the values just used. Those indices could produce the same value-triplet, not a new answer. Do this after recording a match: equal values can be necessary, as in `[0, 0, 0]`.

Deduplicating the anchor prevents repeated triplets across outer-loop iterations; deduplicating `j` and `k` prevents repeats within one anchor. Both are needed because they cover different positions in the triplet.

### Debugging lessons

- Every path through the inner loop must move `j` or `k`, or terminate. A `continue` that leaves both unchanged repeats the same state and can create an infinite loop.
- Initialize `j = i + 1` and `k = len(nums) - 1` once for each new `i`. Then let the inner loop move them. Resetting them inside that loop discards progress and can prevent termination. In Go, `j := i + 1` is fine at the start of the outer-loop iteration; its placement matters.

## Run tests

From the repository root:

```sh
go test ./go/0015-3sum
```
