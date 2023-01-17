package utils

import "golang.org/x/exp/constraints"

// Number constrains numeric values.
type Number interface {
	constraints.Integer | constraints.Float
}

// Signed constrains to Numbers where negative values can be represented.
type Signed interface {
	constraints.Signed | constraints.Float
}

// SignedInt to integers where negative values can be represented.
type SignedInt interface {
	constraints.Signed
}
