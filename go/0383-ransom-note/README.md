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

## Run tests

From the repository root:

```sh
go test ./go/0383-ransom-note
```
