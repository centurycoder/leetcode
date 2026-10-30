# 71. Simplify Path

**Difficulty:** Medium

Given an absolute Unix-style path, return its simplified canonical path. A single `.` is the current directory, `..` moves to the parent directory (without moving above root), and repeated slashes are treated as one. Names containing other sequences of periods, such as `...`, are valid directory names.

**LeetCode Go signature:** `func simplifyPath(path string) string`

## Main idea

Split the path into components and use a slice as a stack of directory names. Ignore empty components and `.`, pop one directory for `..` when the stack is non-empty, and push every other component. Join the remaining names with `/` and prefix the result with `/` to produce the canonical path.

## Examples

- `"/home/"` → `"/home"`
- `"/home//foo/"` → `"/home/foo"`
- `"/home/user/Documents/../Pictures"` → `"/home/user/Pictures"`
- `"/../"` → `"/"`
- `"/.../a/../b/c/../d/./"` → `"/.../b/d"`

## Constraints

- `1 <= len(path) <= 3000`
- `path` is an absolute Unix-style path containing English letters, digits, `.`, `/`, or `_`.

## Run tests

From the repository root:

```sh
go test ./go/0071-simplify-path
```
