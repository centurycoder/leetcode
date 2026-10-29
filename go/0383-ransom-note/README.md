# 383. Ransom Note

**Difficulty:** Easy

Given strings `ransomNote` and `magazine`, return `true` if `ransomNote` can be constructed using letters from `magazine`, or `false` otherwise. Each letter in `magazine` can be used only once.

**LeetCode Go signature:** `func canConstruct(ransomNote string, magazine string) bool`

## Examples

- `ransomNote = "a"`, `magazine = "b"` → `false`
- `ransomNote = "aa"`, `magazine = "ab"` → `false`
- `ransomNote = "aa"`, `magazine = "aab"` → `true`

## Constraints

- `1 <= len(ransomNote), len(magazine) <= 10^5`
- `ransomNote` and `magazine` consist of lowercase English letters

## Implementation summary

Build a `map[rune]int` of the available letter counts from `magazine`. Then scan `ransomNote`, requiring a positive count for each rune and decrementing it when used. A missing or exhausted count means construction is impossible.

For `M = len(magazine)` and `N = len(ransomNote)`, the time complexity is `O(M + N)`. Auxiliary space is `O(U)`, where `U` is the number of distinct letters in `magazine` (at most 26 under the constraints).

## Possible improvements

- Use a `[26]int` and byte indexing instead of `map[rune]int`. The inputs are lowercase ASCII, so this avoids map allocation/lookups and UTF-8 decoding.
- Simplify the map-based check to `if m[c] == 0 { return false }; m[c]--`. A missing map key reads as zero, so this handles both missing and exhausted letters.
- Return `false` immediately when `len(ransomNote) > len(magazine)`. The lowercase ASCII constraint makes byte lengths comparable, and this skips counting for an impossible input.

## Run tests

From the repository root:

```sh
go test ./go/0383-ransom-note
```
