package inplace_test

import (
	"testing"

	"github.com/EduardGomezEscandell/algo/algo"
	"github.com/EduardGomezEscandell/algo/internal/inplace"
	"github.com/EduardGomezEscandell/algo/utils"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	t.Parallel()
	t.Run("int", testMap[int])
	t.Run("int8", testMap[int8])
	t.Run("int32", testMap[int32])
	t.Run("int64", testMap[int64])
}

func TestAdjacentMap(t *testing.T) {
	t.Parallel()
	t.Run("int", testAdjacentMap[int])
	t.Run("int8", testAdjacentMap[int8])
	t.Run("int32", testAdjacentMap[int32])
	t.Run("int64", testAdjacentMap[int64])
}

func TestZipWith(t *testing.T) {
	t.Parallel()
	t.Run("int", testZipWith[int])
	t.Run("int8", testZipWith[int8])
	t.Run("int32", testZipWith[int32])
	t.Run("int64", testZipWith[int64])
}

func testMap[T utils.Signed](t *testing.T) { //nolint: thelper
	t.Parallel()

	testCases := map[string]struct {
		input []T
		op    func(T) int
		want  []int
	}{
		"empty sign":  {input: []T{}, op: algo.Sign[T], want: []int{}},
		"small sign":  {input: []T{1, -2, 3}, op: algo.Sign[T], want: []int{1, -1, 1}},
		"normal sign": {input: []T{-8, 7, 0, 3, 3, -15}, op: algo.Sign[T], want: []int{-1, 1, 0, 1, 1, -1}},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := make([]int, len(tc.input))
			inplace.Map(got, tc.input, tc.op)
			require.Equal(t, tc.want, got)
		})
	}
}

func testZipWith[T utils.Signed](t *testing.T) { //nolint: thelper
	t.Parallel()
	testCases := map[string]struct {
		input1 []T
		input2 []T
		zip    func(T, T) T
		want   []T
	}{
		"empty":      {zip: utils.Sub[T], want: []T{}},
		"half empty": {zip: utils.Sub[T], input1: []T{1}, want: []T{}},
		"single":     {zip: utils.Sub[T], input1: []T{1}, input2: []T{2}, want: []T{-1}},
		"normal":     {zip: utils.Sub[T], input1: []T{1, 3, 9}, input2: []T{2, -1, 6}, want: []T{-1, 4, 3}},
		"normal add": {zip: utils.Add[T], input1: []T{1, 3, 9}, input2: []T{2, -1, 6}, want: []T{3, 2, 15}},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := make([]T, utils.Min(len(tc.input1), len(tc.input2)))
			inplace.ZipWith(got, tc.input1, tc.input2, tc.zip)
			require.Equal(t, tc.want, got)
		})
	}
}

func testAdjacentMap[T utils.Signed](t *testing.T) { //nolint: thelper
	t.Parallel()
	testCases := map[string]struct {
		input   []T
		op      func(T, T) T
		expects []T
	}{
		"empty sum":          {input: []T{}, op: utils.Add[T], expects: []T{}},
		"single sum":         {input: []T{1}, op: utils.Add[T], expects: []T{}},
		"small sum":          {input: []T{1, -2, 3}, op: utils.Add[T], expects: []T{-1, 1}},
		"normal sum":         {input: []T{-8, 7, 0, 3, 3, -15}, op: utils.Add[T], expects: []T{-1, 7, 3, 6, -12}},
		"empty subtraction":  {input: []T{}, op: utils.Sub[T], expects: []T{}},
		"small subtraction":  {input: []T{1, -2, 3}, op: utils.Sub[T], expects: []T{3, -5}},
		"normal subtraction": {input: []T{-8, 7, 0, 3, 3, -15}, op: utils.Sub[T], expects: []T{-15, 7, -3, 0, 18}},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := make([]T, utils.Max(len(tc.input)-1, 0))
			inplace.AdjacentMap(got, tc.input, tc.op)
			require.Equal(t, tc.expects, got)
		})
	}
}
