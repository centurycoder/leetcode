# 125. Valid Palindrome

**Difficulty:** Easy

A phrase is a palindrome when, after converting uppercase letters to lowercase and removing non-alphanumeric characters, it reads the same forward and backward.

**LeetCode Go signature:** `func isPalindrome(s string) bool`

## Examples

- `"A man, a plan, a canal: Panama"` → `true`
- `"race a car"` → `false`
- `" "` → `true`

## Constraints

- `1 <= len(s) <= 2 * 10^5`
- `s` contains printable ASCII characters

## Implementation summary

The function scans inward from both ends of the original string. It skips bytes that are not ASCII letters or digits, compares the remaining bytes after ASCII case folding, and returns `false` on the first mismatch. If the indexes meet or cross without a mismatch, it returns `true`.

`isASCIIAlphanumeric` recognizes ASCII letters and digits. `toLowerASCII` converts uppercase ASCII bytes during comparison, so the function does not allocate a normalized copy.

The scan takes `O(n)` time and `O(1)` extra space.

## Improvement suggestions

- Rename the scan indexes `i` and `j` to `left` and `right` to make their roles clearer.
- Replace the magic number `32` in `toLowerASCII` with `'a' - 'A'` to express the ASCII case offset directly.

## Run tests

From the repository root:

```sh
go test ./go/0125-valid-palindrome
```
