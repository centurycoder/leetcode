package solution

import "strconv"

type Stack struct {
	storage []int
}

func (stack *Stack) push(val int) {
	stack.storage = append(stack.storage, val)
}

func (stack *Stack) pop() int {
	len := len(stack.storage)
	top := stack.storage[len-1]
	stack.storage = stack.storage[:len-1]
	return top
}

func evalRPN(tokens []string) int {
	stack := Stack{}
	result := 0
	for _, token := range tokens {
		switch token {
		case "+":
			num2 := stack.pop()
			num1 := stack.pop()
			result = num1 + num2
			stack.push(result)
		case "-":
			num2 := stack.pop()
			num1 := stack.pop()
			result = num1 - num2
			stack.push(result)
		case "*":
			num2 := stack.pop()
			num1 := stack.pop()
			result = num1 * num2
			stack.push(result)
		case "/":
			num2 := stack.pop()
			num1 := stack.pop()
			result = num1 / num2
			stack.push(result)
		default:
			value, _ := strconv.Atoi(token)
			stack.push(value)
		}
	}
	return stack.pop()
}
