// Package algo implements various classical algorithms.
package algo

import (
	"fmt"
	"sort"

	"github.com/EduardGomezEscandell/algo/internal/inplace"
	"github.com/EduardGomezEscandell/algo/utils"
)

// Map applies function f:T->O element-wise to generate another
// array []O of the same size.
func Map[T, O any](arr []T, f func(T) O) []O {
	if len(arr) < 1 {
		return []O{}
	}

	o := make([]O, len(arr))
	inplace.Map(o, arr, f)
	return o
}

// Foreach applies non-pure function f:T element-wise t modify the array.
func Foreach[T any](arr []T, f func(*T)) {
	for i := 0; i < len(arr); i++ {
		f(&arr[i])
	}
}

// Fill generates an array of length len, where arr[i] = t.
func Fill[T any](len int, t T) []T {
	arr := make([]T, 0, len)
	inplace.Fill(arr, t)
	return arr
}

// Generate generates an array of length len, where arr[i] = f()
// The function will be called in sequential order.
func Generate[T any](len int, f func() T) []T {
	arr := make([]T, 0, len)
	inplace.Generate(arr, f)
	return arr
}

// Generate2D generates a 2D array of lengths n x m, where arr[i][j] = f()
// The function will be called in sequential order.
func Generate2D[T any](n, m int, f func() T) [][]T {
	return Generate(n, func() []T { return Generate(m, f) })
}

// Generate3D generates a 3D array of lengths n x m x p, where arr[i][j][k] = f()
// The function will be called in sequential order.
func Generate3D[T any](n, m, p int, f func() T) [][][]T {
	return Generate(n, func() [][]T { return Generate2D(m, p, f) })
}

// Reduce []T->O applies the function fold:MxT->M cummulatively,
// starting with the default value for M. The end result is
// equivalent to:
//
//	fold(fold(...fold(0, arr[0]), ..., arr[n-2]), arr[n-1])
//
// where n is the length of arr
//
// Example use: Sum the values
//
//	Reduce(arr, func(x,y int)int { return x+y }) # Option 1.
//	Reduce(arr, Add[int])                    # Option 2.
func Reduce[T, O any](arr []T, fold func(O, T) O, init O) O {
	o := init
	for _, a := range arr {
		o = fold(o, a)
	}
	return o
}

// MapReduce maps with the unary operator T->M, producing an
// intermediate array []M that is then reduced with fold:
// OxM->O.
//
// Equivalent to:
//
//	Reduce(Map(arr, unary), fold)
//
// Note: the intermediate array is not stored in memory.
func MapReduce[T, O, M any](arr []T, unary func(T) M, fold func(O, M) O, init O) O {
	o := init
	for _, a := range arr {
		o = fold(o, unary(a))
	}
	return o
}

// ZipWith takes two arrays of type []L and []R, and applies zip:LxR->O
// elementwise to produce an array of type []O and length equal to the
// length of the shortest input.
func ZipWith[L, R, O any](first []L, second []R, f func(L, R) O) []O {
	o := make([]O, utils.Min(len(first), len(second)))
	inplace.ZipWith(o, first, second, f)
	return o
}

// ZipReduce takes two arrays of type []L and []R, and applies zip:LxR->M
// elementwise to produce an intermediate array of type []M and length
// equal to the length of the shortest input. This array is then reduced
// with fold expression fold:OxM->O with initial value 'init'.
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
func ZipReduce[L, R, M, O any](
	first []L,
	second []R,
	zip func(L, R) M,
	fold func(O, M) O,
	init O,
) O {
	ln := utils.Min(len(first), len(second))

	for i := 0; i < ln; i++ {
		init = fold(init, zip(first[i], second[i]))
	}
	return init
}

// AdjacentMap slides a window of size 2 across the array arr applying operator 'f'
// to produce an array of size len(arr)-1.
func AdjacentMap[T, M any](arr []T, f func(T, T) M) []M {
	if len(arr) < 1 {
		return []M{}
	}
	o := make([]M, len(arr)-1)
	inplace.AdjacentMap(o, arr, f)
	return o
}

// AdjacentReduce slides a window of size 2 across the array arr applying operator 'zip',
// producing an intermediate array of size len(arr)-1. This array is then reduced with the
// fold operator.
//
// Equivalent to:
//
//	Reduce(AdjacentMap(arr, zip), fold)
//
// Note: the intermediate array is not stored in memory.
func AdjacentReduce[T, A, I any](arr []T, zip func(T, T) I, fold func(A, I) A) A {
	var acc A
	if len(arr) < 2 {
		return acc
	}

	for i := 1; i < len(arr); i++ {
		acc = fold(acc, zip(arr[i-1], arr[i]))
	}
	return acc
}

// First scans the array and tries to find the first entry according to
// the total ordering defined by the comparator.
//
// Example use: return the maximum value in the list
//
//	First(arr, func(x, y int) bool { return x>y })
//
// Complexity is O(|arr|).
func First[T any](arr []T, comp utils.Comparator[T]) (acc T) {
	if len(arr) == 0 {
		return acc
	}

	acc = arr[0]
	for _, v := range arr[1:] {
		if comp(v, acc) {
			acc = v
		}
	}

	return acc
}

// FirstN returns the first N entries in the array arr according to
// the total ordering defined by the comparator.
//
//	Anticommutativity: isBetter(x,y) = !isBetter(y,x)
//	Total ordering:    isBetter(x,y), isBetter(y,z) <=> isBetter(x, z).
//
// Example use: return the 3 largest values in the list
//
//	FirstN(arr, 3, func(x, y int) bool { return x>y })
//
// Complexity is O(n·|arr|).
func FirstN[T any](arr []T, n uint, comp utils.Comparator[T]) []T {
	if uint(len(arr)) <= n {
		n = uint(len(arr))
		acc := make([]T, n)
		copy(acc, arr[:n])
		Sort(acc, comp)
		return acc
	}
	acc := make([]T, n)
	copy(acc, arr[:n])
	Sort(acc, comp)

	for _, a := range arr[n:] {
		InsertSorted(n, acc, a, comp)
	}

	return acc
}

// InsertSorted takes a list arr sorted according to comp and emplaces
// x in the position that keeps the list sorted. After this, the last
// element is dropped from the list. Note that if x were to be last,
// the emplace-then-drop step is optimized out.
//
// Complexity is O(n).
func InsertSorted[T any](n uint, arr []T, x T, comp utils.Comparator[T]) {
	if comp(arr[n-1], x) { // Not top n
		return
	}

	i := int(n) - 1
	for ; i > 0; i-- {
		if comp(arr[i-1], x) {
			arr[i] = x
			break
		}
		arr[i] = arr[i-1]
	}
	arr[i] = x
}

// Sort sorts a list according to a comparator comp. Item i preceedes
// item j <=> comp(i,j) is true.
//
// Example: sort from smallest to largest:
//
//	Sort(arr, func(l,r int) bool { return l<r }) // Sorts incrementally
//	Sort(arr, utils.Lt)                          // The same, but shorter
//
// Complexity is O(|arr|·log(|arr|)).
func Sort[T any](arr []T, comp utils.Comparator[T]) {
	sort.Slice(arr, func(i, j int) bool {
		return comp(arr[i], arr[j])
	})
}

// Intersect finds all elements that two slices have in common.
//
// The lists are expected to have been sorted with the comparator 'comp'.
// Two items a,b are considered equivalent if both comp(a,b) and comp(b,a)
// are false.
//
// It returns a slice with their common items. Items in the output list are
// repeated as many times as the smallest number of repetitions between the
// two lists.
//
// Complexity is O(|first| + |second|).
func Intersect[T any](first, second []T, comp utils.Comparator[T]) []T {
	common := []T{}
	var f, s int
	for f < len(first) && s < len(second) {
		// first[f] preceedes second[s]
		if comp(first[f], second[s]) {
			f++
			continue
		}
		// first[f] succeeds second[s]
		if comp(second[s], first[f]) {
			s++
			continue
		}
		// first[f] == second[s]
		common = append(common, first[f])
		f++
		s++
	}
	return common
}

// Unique modifies array arr so that all unique items are moved to the
// beginning. Returns the index where the new end is.
func Unique[T any](arr []T, equal utils.Comparator[T]) (endUnique int) {
	if len(arr) == 0 {
		return 0
	}

	endUnique = 1
	for i := 1; i < len(arr); i++ {
		if equal(arr[i], arr[endUnique-1]) {
			continue
		}
		if i != endUnique {
			arr[endUnique], arr[i] = arr[i], arr[endUnique] // Swap
		}
		endUnique++
	}

	return endUnique
}

// Reverse returns a a copy of the original slice with its elements in opposite order.
func Reverse[T any](arr []T) []T {
	out := make([]T, 0, len(arr))
	for j := len(arr) - 1; j >= 0; j-- {
		out = append(out, arr[j])
	}
	return out
}

// Stride takes one value every n of them, stores it in the output array and drops the rest.
func Stride[T any](in []T, n int) []T {
	out := make([]T, 0, len(in)/3)
	for i := 0; i < len(in); i += n {
		out = append(out, in[i])
	}
	return out
}

// Find traverses array arr searching for an element that matches val
// according to comparator eq and returs its index. If none match,
// -1 is returned.
func Find[T any](arr []T, val T, eq utils.Comparator[T]) int {
	return FindIf(arr, func(t T) bool { return eq(t, val) })
}

// FindIf traverses array arr searching for an element that makes
// f return true, and returs its index. If none match, -1 is returned.
func FindIf[T any](arr []T, pred utils.Predicate[T]) int {
	for i, v := range arr {
		if pred(v) {
			return i
		}
	}
	return -1
}

// Partition rearranges a list such that
//
//	pred(arr[i]) is true <=> i < j
//
// and returns this value j. If the entire list fulfils the predicate,
// j will be equal to the length.
//
// Example: partition smaller than 3:
//
//		j := Partition(arr, func(x int) bool { return l<3 })
//
//	 arr -> [numbers, less, than 3, numbers, greater, than 3]
//	                                ^
//	                                j
//
// Complexity is O(|arr|).
func Partition[T any](slice []T, predicate func(t T) bool) (p int) {
	if len(slice) == 0 {
		return 0
	}
	for i := range slice {
		if !predicate(slice[i]) {
			continue
		}
		if i == p {
			p++
			continue
		}
		slice[i], slice[p] = slice[p], slice[i]
		p++
	}
	return p
}

// Insert returns an array {arr[:position], value, arr[position:]}.
// Original array becomes invalidated. Usage:
//
//	arr = Insert(arr, "hello", 5)
func Insert[T any](arr []T, value T, position int) []T {
	if position < 0 || position > len(arr) {
		panic(fmt.Errorf("index %d out of range [0, %d)", position, len(arr)+1))
	}
	var t T
	arr = append(arr, t) // Dummy entry, will be overwritten
	for i := len(arr) - 1; i > position; i-- {
		arr[i] = arr[i-1]
	}
	arr[position] = value
	return arr
}

// Rotate rotates an array to the left when n>0, and to the right
// otherwise; such that all elements are shifted left/right by n
// positions. Items that would be shifted out of the range are
// shifted back in from the other side.
//
// Returns the position where the former first item is now located.
func Rotate[T any](arr []T, n int) int {
	if Abs(n) >= len(arr) {
		panic(fmt.Errorf("abs(n) must be less than the length of the array (n=%d, len=%d)", n, len(arr)))
	}
	switch {
	case n < 0:
		return rotateRight(arr, uint(-n))
	case n > 0:
		return rotateLeft(arr, uint(n))
	case n == 0:
	}
	return 0
}

func rotateLeft[T any](arr []T, n uint) int {
	aux := make([]T, n)
	k := int(uint(len(arr)) - n)
	copy(aux, arr[:n])
	copy(arr[:k], arr[n:])
	copy(arr[k:], aux)
	return k
}

func rotateRight[T any](arr []T, n uint) int {
	k := int(uint(len(arr)) - n)
	aux := make([]T, n)
	copy(aux, arr[k:])
	copy(arr[n:], arr[:k])
	copy(arr[:n], aux)
	return int(n)
}
