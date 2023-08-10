package leetCode

/**
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/

var typeStack []rune

func push(r rune) {
	if typeStack == nil {
		typeStack = make([]rune, 0)
	}
	typeStack = append(typeStack, r)
}

func pop() rune {
	n := len(typeStack)
	if n == 0 {
		return -1
	}

	r := typeStack[n-1]
	typeStack = typeStack[:n-1]
	return r
}

func IsBalancedParenTypes(s string) bool {

	for _, r := range []rune(s) {
		switch r {
		case '(', '{', '[':
			push(r)
		case ')':
			r := pop()
			if r != '(' {
				return false
			}
		case '}':
			r := pop()
			if r != '{' {
				return false
			}
		case ']':
			r := pop()
			if r != '[' {
				return false
			}
		}
	}

	return true
}
