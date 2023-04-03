package generics

import "fmt"

type Printable interface {
	String() string
}

type Stack []Printable

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(v Printable) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (Printable, error) {
	if s.isEmpty() {
		return nil, fmt.Errorf("cannot pop, stack is empty")
	}

	lastIdx := len(*s) - 1
	last := (*s)[lastIdx]
	(*s) = (*s)[:lastIdx]
	return last, nil
}
