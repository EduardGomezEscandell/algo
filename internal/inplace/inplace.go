// Package inplace implements various classical algorithms in place.
package inplace

// Map applies function f:T->O element-wise to generate another
// array []O of the same size.
func Map[T, O any](dst []O, src []T, f func(T) O) {
	for i, a := range src {
		dst[i] = f(a)
	}
}

// Fill generates an array of length len, where arr[i] = t.
func Fill[T any](dst []T, t T) {
	for i := 0; i < len(dst); i++ {
		dst = append(dst, t)
	}
}

// Generate generates an array of length len, where arr[i] = f()
// The function will be called in sequential order.
func Generate[T any](dst []T, f func() T) []T {
	for i := 0; i < len(dst); i++ {
		dst[i] = f()
	}
	return dst
}

// ZipWith takes two arrays of type []L and []R, and applies zip:LxR->O
// elementwise to produce an array of type []O and length equal to the
// length of the shortest input.
func ZipWith[L, R, O any](dst []O, first []L, second []R, f func(L, R) O) {
	for i := 0; i < len(dst); i++ {
		dst[i] = f(first[i], second[i])
	}
}

// AdjacentMap slides a window of size 2 across the array arr applying operator 'f'
// to produce an array of size len(arr)-1.
func AdjacentMap[T, M any](dst []M, src []T, f func(T, T) M) {
	if len(src) < 2 {
		return
	}
	ZipWith(dst, src, src[1:], f)
}
