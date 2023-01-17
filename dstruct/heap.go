// Package dstruct contains data stuctures.
package dstruct

import (
	"container/heap"

	"github.com/EduardGomezEscandell/algo/utils"
)

// Heap implemements a heap data structure
// by wrapping around the standard library's heap.
type Heap[T any] struct {
	impl *heapImpl[T]
}

// NewHeap creates and initializes a heap.
func NewHeap[T any](best utils.Comparator[T]) Heap[T] {
	return HeapFromSlice([]T{}, best)
}

// HeapFromSlice creates and initializes a heap from a given slice.
func HeapFromSlice[T any](src []T, best utils.Comparator[T]) Heap[T] {
	h := Heap[T]{
		&heapImpl[T]{
			data: src,
			comp: best,
		},
	}
	heap.Init(h.impl)
	return h
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Push(t T) {
	heap.Push(h.impl, t)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (h *Heap[T]) Pop() T {
	return heap.Pop(h.impl).(T) //nolint: forcetypeassert
}

// Len returns the size of the heap.
func (h Heap[T]) Len() int {
	return h.impl.Len()
}

// Data returns a pointer to the internal data
// Run Heap.Fix if you modify it in any way.
func (h Heap[T]) Data() *[]T {
	return &h.impl.data
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (h Heap[T]) Remove(i int) T {
	return heap.Remove(h.impl, i).(T) //nolint: forcetypeassert
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func (h Heap[T]) Fix(i int) {
	heap.Fix(h.impl, i)
}

// Repair establishes the heap invariants required by the other routines in this package.
// Repair is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func (h Heap[T]) Repair() {
	heap.Init(h.impl)
}

// implementation

type heapImpl[T any] struct {
	data []T
	comp func(x, y T) bool
}

func (h heapImpl[T]) Len() int           { return len(h.data) }
func (h heapImpl[T]) Less(i, j int) bool { return h.comp(h.data[i], h.data[j]) }
func (h heapImpl[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

// Push emplaces the leading item. DO NOT USE!
// Use heap.Push(h, x) instead.
func (h *heapImpl[T]) Push(x any) {
	h.data = append(h.data, x.(T)) //nolint: forcetypeassert
	// We simply allow a panic ^.
}

// Pop extracts the leading item. DO NOT USE!
// Use heap.Pop(h).(T) instead.
func (h *heapImpl[T]) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[:n-1]
	return x
}
