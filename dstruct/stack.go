// Package dstruct implements various Data STRUCTures.
package dstruct

import (
	"github.com/EduardGomezEscandell/algo/algo"
)

// Stack data structure. Implements a LIFO queue.
type Stack[T any] struct {
	data []T
}

// NewStack creates a new stack, with optional arguments
// for initial size (with default initialization) and
// initial capacity.
func NewStack[T any](args ...int) Stack[T] {
	var size, capacity int
	switch len(args) {
	default:
		panic("Only two args allowed: size and capacity")
	case 2:
		capacity = args[1]
		fallthrough
	case 1:
		size = args[0]
		fallthrough
	case 0:
		return Stack[T]{data: make([]T, size, capacity)}
	}
}

// Size is the count of elements in the stack.
func (s Stack[T]) Size() int {
	return len(s.data)
}

// IsEmpty indicates when the count of elements in the stack is zero.
func (s Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Peek reveals a copy of the item at the top of the stack.
func (s Stack[T]) Peek() T {
	if s.IsEmpty() {
		panic("peeking into empty stack")
	}
	return s.data[len(s.data)-1]
}

// Pop removes an item to the top of the stack.
func (s *Stack[T]) Pop() {
	if s.IsEmpty() {
		panic("peeking into empty stack")
	}
	s.data = s.data[:len(s.data)-1]
}

// Push inserts an item to the top of the stack.
func (s *Stack[T]) Push(t T) {
	s.data = append(s.data, t)
}

// Data reveals the internal storage. The storage is
// arranged from bottom to top of the stack.
func (s *Stack[T]) Data() []T {
	return s.data
}

// Invert reverses the order of the items within the stack.
func (s *Stack[T]) Invert() {
	s.data = algo.Reverse(s.data)
}
