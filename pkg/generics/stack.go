package generics

import "fmt"

type Stack []any

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(v any) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (any, error) {
	if s.isEmpty() {
		return nil, fmt.Errorf("cannot pop, stack is empty")
	}

	lastIdx := len(*s) - 1
	last := (*s)[lastIdx]
	(*s) = (*s)[:lastIdx]
	return last, nil
}
