package algo

import (
	"github.com/EduardGomezEscandell/algo/utils"
	"golang.org/x/exp/constraints"
)

// GCD computes the greatest common divisor (GCD) via Euclidean algorithm.
func GCD[T constraints.Integer](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM finds the Least Common Multiple (LCM) via GCD.
func LCM[T constraints.Integer](a, b T, integers ...T) T {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// Sign returns +1 or -1, with the same sign as `a`.
// Returns 0 if a is 0.
func Sign[T utils.Signed](a T) int {
	switch {
	case a > 0:
		return 1
	case a < 0:
		return -1
	}
	return 0
}

// Abs returns the absolute value of a, i.e. a scalar
// with the same magnitude and with positive sign.
func Abs[T utils.Signed](a T) T {
	return T(Sign(a)) * a
}

// Count adds one to `a` iff `b` is true.
func Count[T utils.Number](a T, b bool) T {
	if b {
		return a + 1
	}
	return a
}

// Clamp returns value x if it is inside the range [lo, hi], otherwise
// returning the range boundary it is closest to.
func Clamp[T utils.Number](lo, x, hi T) T {
	return utils.Max(lo, utils.Min(x, hi))
}
