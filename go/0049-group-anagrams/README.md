# 49. Group Anagrams

**Difficulty:** Medium

Given a slice of strings, group together strings that are anagrams. The order of the groups and the strings within each group does not matter.

**LeetCode Go signature:** `func groupAnagrams(strs []string) [][]string`

## Examples

- `["eat", "tea", "tan", "ate", "nat", "bat"]` → groups containing `["eat", "tea", "ate"]`, `["tan", "nat"]`, and `["bat"]`
- `[""]` → `[[""]]`
- `["a"]` → `[["a"]]`

## Constraints

- `1 <= len(strs) <= 10^4`
- `0 <= len(strs[i]) <= 100`
- Each string contains only lowercase English letters.

## Key idea: use a frequency array as the map key

Anagrams contain the same count of every letter. Build a 26-element array where each position counts one lowercase letter: index `0` counts `a`, index `1` counts `b`, and so on. Use that array as the map key:

```go
var key [26]int
for i := 0; i < len(s); i++ {
    key[s[i]-'a']++
}
groups[key] = append(groups[key], s)
```

For example, `"eat"` and `"tea"` both have one `a`, one `e`, and one `t`, so their arrays are equal. `"ac"` has one `a` and one `c`, while `"bb"` has two `b`s, so their arrays differ.

Go arrays are comparable and can be map keys. A slice such as `[]int` cannot be a map key because slices are not comparable. The fixed-size array also avoids sorting each string to construct its key.

## First-version issue

The first version used an integer hash/key, but distinct letter combinations such as `"ac"` and `"bb"` could produce the same integer. If two strings are stored under the same integer key, the map treats them as the same group; it cannot recover the different letter counts. The frequency array preserves the full count signature, so those combinations have different keys.

## Run tests

From the repository root:

```sh
go test ./go/0049-group-anagrams
```
