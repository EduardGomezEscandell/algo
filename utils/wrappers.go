package utils

// Add is a wrapper around the plus operator.
func Add[T Number](a, b T) T {
	return a + b
}

// Sub is a wrapper around the minus operator.
func Sub[T Number](a, b T) T {
	return a - b
}

// Mul is a wrapper around the product operator.
func Mul[T Number](a, b T) T {
	return a * b
}

// Div is a wrapper around the division operator.
// Unsafe from division by zero.
func Div[T Number](a, b T) T {
	return a / b
}

// Min returns the minimum between two numbers.
func Min[T Number](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

// Max returns the maximum between two numbers.
func Max[T Number](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

// Gt is wrapper around the > operator.
func Gt[T Number](a, b T) bool {
	return a > b
}

// Ge is wrapper around the >= operator.
func Ge[T Number](a, b T) bool {
	return a >= b
}

// Eq is wrapper around the == operator.
func Eq[T Number](a, b T) bool {
	return a == b
}

// Le is wrapper around the <= operator.
func Le[T Number](a, b T) bool {
	return a <= b
}

// Lt is wrapper around the < operator.
func Lt[T Number](a, b T) bool {
	return a < b
}

// And is the logical or.
func And(a, b bool) bool {
	return a && b
}

// Or is the logical or.
func Or(a, b bool) bool {
	return a || b
}

// Append is a wrapper around the append built-in.
func Append[T any](acc []T, t T) []T {
	return append(acc, t)
}
