// Package palgo implements parallel versions various classical algorithms.
package palgo

import (
	"github.com/EduardGomezEscandell/algo/algo"
	"github.com/EduardGomezEscandell/algo/utils"
)

// Map applies function f:T->O element-wise to generate another
// array []O of the same size.
func Map[T, O any](arr []T, f func(T) O) []O {
	o := make([]O, len(arr))
	work := workAllocation(len(arr))
	if len(work) < 2 {
		return algo.Map(arr, f)
	}
	distribute(work, func(w workAlloc) {
		for i := w.begin; i < w.end; i++ {
			o[i] = f(arr[i])
		}
	})

	return o
}

// Foreach applies non-pure function f:T element-wise t modify the array.
func Foreach[T any](arr []T, f func(*T)) {
	work := workAllocation(len(arr))
	if len(work) < 2 {
		algo.Foreach(arr, f)
		return
	}
	distribute(work, func(w workAlloc) {
		algo.Foreach(arr[w.begin:w.end], f)
	})
}

// Fill generates an array of length len, where arr[i] = t.
func Fill[T any](n int, t T) []T {
	work := workAllocation(n)
	if len(work) < 2 {
		return algo.Fill(n, t)
	}
	o := make([]T, n)
	distribute(work, func(w workAlloc) {
		for it := w.begin; it != w.end; it++ {
			o[it] = t
		}
	})
	return o
}

// Reduce []T->O applies the function fold:TxT->T cummulatively,
// starting with the initial value init. The end result is
// equivalent to:
//
//	fold(fold(...fold(0, arr[0]), ..., arr[n-2]), arr[n-1])
//
// where n is the length of arr.
// The fold must be associative.
//
// Example use: Sum the values
//
//	Reduce(arr, func(x,y int)int { return x+y }) # Option 1.
//	Reduce(arr, Add[int])                        # Option 2.
func Reduce[T any](arr []T, fold func(T, T) T, init T) T {
	work := workAllocation(len(arr))
	if len(work) < 2 {
		return algo.Reduce(arr, fold, init)
	}

	o := make([]T, len(work))
	distribute(work, func(w workAlloc) {
		init := fold(arr[w.begin], arr[w.begin+1])
		w.begin += 2
		o[w.worker] = algo.Reduce(arr[w.begin:w.end], fold, init)
	})
	return algo.Reduce(o, fold, init)
}

// MapReduce maps with the unary operator unary:T->O, producing an
// intermediate array []O that is then reduced with an associative
// fold:OxO->O.
//
// Equivalent to:
//
//	Reduce(Map(arr, unary), fold)
//
// Note: the intermediate array is not stored in memory.
func MapReduce[T, O any](arr []T, unary func(T) O, fold func(O, O) O, init O) O {
	work := workAllocation(len(arr))
	if len(work) < 2 {
		return algo.MapReduce(arr, unary, fold, init)
	}

	o := make([]O, len(work))
	distribute(work, func(w workAlloc) {
		init := fold(unary(arr[w.begin]), unary(arr[w.begin+1]))
		w.begin += 2
		o[w.worker] = algo.MapReduce(arr[w.begin:w.end], unary, fold, init)
	})
	return algo.Reduce(o, fold, init)
}

// ZipWith takes two arrays of type []L and []R, and applies zip:LxR->O
// elementwise to produce an array of type []O and length equal to the
// length of the shortest input.
func ZipWith[L, R, O any](first []L, second []R, f func(L, R) O) []O {
	ln := utils.Min(len(first), len(second))
	work := workAllocation(ln)
	if len(work) < 2 {
		return algo.ZipWith(first, second, f)
	}

	o := make([]O, ln)
	distribute(work, func(w workAlloc) {
		for i := w.begin; i < w.end; i++ {
			o[i] = f(first[i], second[i])
		}
	})
	return o
}

// ZipReduce takes two arrays of type []L and []R, and applies zip:LxR->M
// elementwise to produce an intermediate array of type []O and length
// equal to the length of the shortest input. This array is then reduced
// with fold expression fold:OxO->O with initial value 'init'.
//
// Equivalent to:
//
//	Reduce(ZipWith(first, second, zip), fold)
//
// Note: the intermediate array is not stored in memory.
//
// Example: compute the inner product (u, v):
//
//	ZipReduce(u, v, utils.Mul, utils.Add, 0)
func ZipReduce[L, R, O any](
	first []L,
	second []R,
	zip func(L, R) O,
	fold func(O, O) O,
	init O,
) O {
	ln := utils.Min(len(first), len(second))
	work := workAllocation(ln)
	if len(work) < 2 {
		return algo.ZipReduce(first, second, zip, fold, init)
	}

	o := make([]O, len(work))
	distribute(work, func(w workAlloc) {
		a := zip(first[w.begin], second[w.begin])
		w.begin++
		b := zip(first[w.begin], second[w.begin])
		w.begin++
		o[w.worker] = algo.ZipReduce(first[w.begin:w.end], second[w.begin:w.end], zip, fold, fold(a, b))
	})

	return algo.Reduce(o, fold, init)
}
