package main

import (
	"fmt"
	"sync"
)

// Stack
type Stack struct {
	lock sync.RWMutex
	data []interface{}
}

func NewStack(args ...int) *Stack {
	defaultSize := 1 << 4
	if len(args) > 0 {
		defaultSize = args[0]
	}

	return &Stack{
		data: make([]interface{}, 0, defaultSize),
	}
}

func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = append(s.data, item)
}

func (s *Stack) Pop() interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if s.IsEmpty() {
		return nil
	}
	index := len(s.data) - 1
	item := s.data[index]
	s.data = s.data[:index]

	return item
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Size() int {
	return len(s.data)
}

func main() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
