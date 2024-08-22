package stack

import "fmt"

const SIZE_8MB = 8 * 1024 * 1024

type Stack struct {
	store [][]byte
}

func NewStack(size int) *Stack {
	store := make([][]byte, size)

	for index := range store {
		store[index] = make([]byte, SIZE_8MB)
	}

	return &Stack{
		store: store,
	}
}

func (s *Stack) Pop() ([]byte, error) {
	if len(s.store) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	last := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]

	return last, nil
}

func (s *Stack) Push(data []byte) {
	s.store = append(s.store, data)
}
