# 392. Is Subsequence

**Difficulty:** Easy

Given strings `s` and `t`, return whether `s` is a subsequence of `t`. A subsequence keeps the characters in their original order but does not need to be contiguous.

**LeetCode Go signature:** `func isSubsequence(s string, t string) bool`

## Examples

- `s = "abc"`, `t = "ahbgdc"` → `true`
- `s = "axc"`, `t = "ahbgdc"` → `false`
- `s = ""`, `t = "ahbgdc"` → `true`

## Constraints

- `0 <= len(s) <= 100`
- `0 <= len(t) <= 10^4`
- `s` and `t` contain only lowercase English letters

## Implementation summary

Use `i` to track the next character to match in `s` and `j` to scan `t` from left to right. When `s[i]` matches `t[j]`, advance both indexes; otherwise, advance only `j`. Each match consumes the next needed character in order. The earliest possible match is always safe because it leaves the largest remaining suffix of `t` for later characters in `s`.

The scan succeeds when `i` reaches `len(s)`. This also handles an empty `s`: it is a subsequence of every `t`.

Each byte of `t` is visited at most once, so the time complexity is `O(len(t))`, with `O(1)` extra space.

## Run tests

From the repository root:

```sh
go test ./go/0392-is-subsequence
```
