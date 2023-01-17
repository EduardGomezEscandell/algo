// Package utils has various utils used by algo and dstruct.
package utils

type (
	// Comparator is a function that returns a boolean based on two inputs.
	// It should be stateless, and allow for a total ordering of elements in
	// T such that a precedes b if, and only if, Comparator(a,b)==true.
	// Functions that require a comparator input will assume as much.
	Comparator[T any] func(a, b T) bool

	// Predicate is a function that returns a boolean based on a single input.
	// Furthermore, it should be stateless. Functions that require a predicate
	// input will assume as much.
	Predicate[T any] func(a T) bool
)

// Equal takes returns an equality check. Two items are considered equal
// if their position is equivalent according to the ordering defined by
// the comparator.
func Equal[T any](comp Comparator[T]) Comparator[T] {
	return func(a, b T) bool {
		return !comp(a, b) && !comp(b, a)
	}
}

// Identity does not mutate the object.
func Identity[T any](a T) T {
	return a
}
