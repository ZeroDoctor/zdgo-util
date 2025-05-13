package zdutil

import (
	"fmt"
	"sync"
)

type Stack[T any] struct {
	slice []T
	m     sync.Mutex
}

// NewStack creates a new Stack with an initial value of the given slice.
// If the provided slice is empty, the stack will be empty.
func NewStack[T any](slice ...T) Stack[T] {
	return Stack[T]{slice: slice}
}

// Clear removes all elements from the stack, leaving it empty.
func (s *Stack[T]) Clear() {
	s.m.Lock()
	defer s.m.Unlock()

	s.slice = []T{}
}

// Peek returns the top element of the stack without removing it.
// The function assumes that the stack is not empty.
// It acquires a lock to ensure thread-safe access.
func (s *Stack[T]) Peek() T {
	s.m.Lock()
	defer s.m.Unlock()

	return s.slice[0]
}

// Push adds one or more elements to the top of the stack.
// The function assumes that the stack is not empty.
// It acquires a lock to ensure thread-safe access.
func (s *Stack[T]) Push(elem ...T) {
	s.m.Lock()
	defer s.m.Unlock()

	var arr []T
	arr = append(arr, elem...)
	arr = append(arr, s.slice...)
	s.slice = arr
}

// Pop removes and returns the top element of the stack.
// The function returns nil if the stack is empty.
// It acquires a lock to ensure thread-safe access.
func (s *Stack[T]) Pop() *T {
	s.m.Lock()
	defer s.m.Unlock()

	if len(s.slice) <= 0 {
		return nil
	}

	elem := s.slice[0]
	s.slice = s.slice[1:]

	return &elem
}

// Len returns the number of elements in the stack.
// The function acquires a lock to ensure thread-safe access.
func (s *Stack[T]) Len() int {
	s.m.Lock()
	defer s.m.Unlock()

	return len(s.slice)
}

// String returns a string representation of the stack.
// The string is formatted as "[x,y,z]" where x, y, and z are the elements of the stack.
// The function acquires a lock to ensure thread-safe access.
func (s *Stack[T]) String() string {
	s.m.Lock()
	defer s.m.Unlock()

	if len(s.slice) <= 0 {
		return "[]"
	}

	var str string
	str = "["
	for _, s := range s.slice {
		str += fmt.Sprint(s, ",")
	}
	str = str[:len(str)-1]
	str += "]"
	return str
}
