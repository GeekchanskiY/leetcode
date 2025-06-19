package main

import (
	"fmt"
)

// Time, space: O(n)
func isValid(s string) bool {
	stack := make([]rune, 0, len(s))

	for _, c := range s {
		switch c {
		case '(':
			stack = append(stack, c)
		case ')':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != '(' {
				return false
			}

			stack = stack[:len(stack)-1]
		case '{':
			stack = append(stack, c)
		case '}':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != '{' {
				return false
			}

			stack = stack[:len(stack)-1]
		case '[':
			stack = append(stack, c)
		case ']':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != '[' {
				return false
			}

			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}
	fmt.Println(stack)
	if len(stack) != 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(isValid(""))
	fmt.Println(isValid("[({})]"))
	fmt.Println(isValid("{()[]}"))
	fmt.Println(isValid("{[)}"))
	fmt.Println(isValid("["))
}
