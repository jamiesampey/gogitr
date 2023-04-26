package generics

import "fmt"

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.isEmpty() {
		var zeroValue T
		return zeroValue, fmt.Errorf("cannot pop, stack is empty")
	}

	lastIdx := len(s.data) - 1
	last := s.data[lastIdx]
	s.data = s.data[:lastIdx]

	return last, nil
}

func (s *Stack[T]) PopAll() {
	for {
		if item, err := s.Pop(); err == nil {
			fmt.Printf("Popped item is %s, stack size is %d\n", item, len(s.data))
		} else {
			fmt.Printf("Error is \"%s\"\n", err)
			break
		}
	}
}
