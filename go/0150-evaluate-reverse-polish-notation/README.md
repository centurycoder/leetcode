# 150. Evaluate Reverse Polish Notation

**Difficulty:** Medium

Given tokens representing a valid Reverse Polish Notation expression, evaluate it and return the integer result. The operators are `+`, `-`, `*`, and `/`; integer division truncates toward zero.

**LeetCode Go API:** `func evalRPN(tokens []string) int`

## Implementation summary

The solution uses a `Stack` struct containing `[]int` storage. Integer tokens are parsed with `strconv.Atoi` and pushed. For an operator, it pops the right operand first and then the left operand, applies the operation in left-to-right order, and pushes the result. This operand order is important for subtraction and division. Go's integer division truncates toward zero, matching the problem requirement.

Each token is processed once. Time complexity is O(n); auxiliary space is O(n), where n is the number of tokens.

## Reference implementation

This version uses a slice local to `evalRPN` rather than a stack struct. The `pop` closure captures the local slice variable, so shortening the slice updates the variable used by the rest of the function.

```go
func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	pop := func() int {
		top := len(stack) - 1
		value := stack[top]
		stack = stack[:top]
		return value
	}

	for _, token := range tokens {
		switch token {
		case "+":
			right := pop()
			left := pop()
			stack = append(stack, left+right)
		case "-":
			right := pop()
			left := pop()
			stack = append(stack, left-right)
		case "*":
			right := pop()
			left := pop()
			stack = append(stack, left*right)
		case "/":
			right := pop()
			left := pop()
			stack = append(stack, left/right)
		default:
			value, _ := strconv.Atoi(token)
			stack = append(stack, value)
		}
	}

	return pop()
}
```

Add `import "strconv"` when using this function in a Go source file.

## Why a slice helper may not update the caller

A Go slice is a descriptor containing a pointer to an underlying array, a length, and a capacity. Passing a slice to a function copies that descriptor. Reslicing a parameter changes only the copied length, and `append` may allocate a new backing array and change only the copied descriptor. Therefore, helpers like these do not update the caller's slice header:

```go
func push(stack []int, value int) {
	stack = append(stack, value)
}

func pop(stack []int) int {
	value := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return value
}
```

Use one of these approaches instead:

- Return the updated slice from `push`, and return both the popped value and shortened slice from `pop`; assign those return values at every call site.
- Pass a `*[]int` and update `*stack` inside each helper.
- Keep the slice in a struct and use pointer-receiver methods, as the current implementation does. Updating the struct's slice field also updates its length/capacity descriptor.

The reference implementation avoids separate slice helpers: its closure operates on the `stack` variable in `evalRPN` directly.

## Constraints

- `1 <= tokens.length <= 10^4`
- Each token is one of the four operators or an integer in `[-200, 200]`.
- The expression is valid, contains no division by zero, and all intermediate values fit in a signed 32-bit integer.

## Run tests

From the repository root:

```sh
go test ./go/0150-evaluate-reverse-polish-notation
```
