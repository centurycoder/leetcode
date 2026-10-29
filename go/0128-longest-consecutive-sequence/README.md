# 128. Longest Consecutive Sequence

**Difficulty:** Medium

Given an unsorted integer slice, return the length of its longest sequence of consecutive values. The sequence does not need to be contiguous in the input.

**LeetCode Go signature:** `func longestConsecutive(nums []int) int`

## Key idea: use a set and only scan from sequence starts

Put the input values in a hash set. For each distinct value `num`, begin scanning forward only when `num-1` is absent. That identifies the start of a consecutive run. Count forward until the next value is absent, then update the longest length.

```go
for num := range lookup {
    if _, exists := lookup[num-1]; exists {
        continue
    }

    length := 1
    for next := num + 1; ; next++ {
        if _, exists := lookup[next]; !exists {
            break
        }
        length++
    }
    if length > longest {
        longest = length
    }
}
```

The order of map iteration does not matter because only the maximum length is returned. With average constant-time map lookups, building the set and scanning the runs takes **O(n)** time and **O(n)** space.

## Mistake that caused TLE

The first scan loop iterated over the original `nums` slice to find sequence starts. Duplicate values can then start the same forward scan repeatedly. For example, many copies of `0` followed by `1, 2, ..., k` make every `0` rescan the same run, resulting in quadratic work.

Iterate over `lookup` instead: a map contains each distinct number once. Also, keep the `num-1` check; without it, every value in a long run could rescan the rest of that run.

The allowed values range from `-10^9` to `10^9`, so a direct array indexed by the input value is unsuitable: values can be negative, and the range is too large. A map handles both negative and large values without allocating by numeric range.

## Debugging notes

- The range-limit test expects `5`: `999_999_996` through `1_000_000_000` is a five-number consecutive run.
- During development, that test reported `4` instead of `5`; this was a correctness mismatch separate from the duplicate-rescan TLE.
- An empty input should return `0`.
- Duplicate input values do not extend a consecutive sequence.

## Run tests

From the repository root:

```sh
go test ./go/0128-longest-consecutive-sequence
```
