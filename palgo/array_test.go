package palgo_test

import (
	"testing"

	"github.com/EduardGomezEscandell/algo/algo"
	"github.com/EduardGomezEscandell/algo/palgo"
	"github.com/EduardGomezEscandell/algo/utils"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	t.Parallel()
	t.Run("int", testMap[int])
	// t.Run("int8", testMap[int8])
	// t.Run("int32", testMap[int32])
	// t.Run("int64", testMap[int64])
}

func TestReduce(t *testing.T) {
	t.Parallel()
	t.Run("int", testReduce[int])
	t.Run("int8", testReduce[int8])
	t.Run("int32", testReduce[int32])
	t.Run("int64", testReduce[int64])
}

func TestMapReduce(t *testing.T) {
	t.Parallel()
	t.Run("int", testMapReduce[int])
	t.Run("int8", testMapReduce[int8])
	t.Run("int32", testMapReduce[int32])
	t.Run("int64", testMapReduce[int64])
}

func TestZipWith(t *testing.T) {
	t.Parallel()
	t.Run("int", testZipWith[int])
	t.Run("int8", testZipWith[int8])
	t.Run("int32", testZipWith[int32])
	t.Run("int64", testZipWith[int64])
}

func TestZipReduce(t *testing.T) {
	t.Parallel()
	t.Run("int", testZipReduce[int])
	t.Run("int8", testZipReduce[int8])
	t.Run("int32", testZipReduce[int32])
	t.Run("int64", testZipReduce[int64])
}

func testMap[T utils.Signed](t *testing.T) { //nolint: thelper
	t.Parallel()

	testCases := map[string]struct {
		input   []T
		op      func(T) int
		expects []int
	}{
		"empty sign":  {input: []T{}, op: algo.Sign[T], expects: []int{}},
		"small sign":  {input: []T{1, -2, 3}, op: algo.Sign[T], expects: []int{1, -1, 1}},
		"normal sign": {input: []T{-8, 7, 0, 3, 3, -15}, op: algo.Sign[T], expects: []int{-1, 1, 0, 1, 1, -1}},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := palgo.Map(tc.input, tc.op)
			require.Equal(t, tc.expects, got)
		})
	}
}

func testReduce[T utils.Signed](t *testing.T) { //nolint: thelper // nolint: thelper
	t.Parallel()

	testCases := map[string]struct {
		input   []T
		fold    func(T, T) T
		expects T
	}{
		"empty sum":        {input: []T{}, fold: utils.Add[T], expects: 0},
		"small sum":        {input: []T{1, 2, 3}, fold: utils.Add[T], expects: 6},
		"normal sum":       {input: []T{8, 7, 5, 3, 3, -15}, fold: utils.Add[T], expects: 11},
		"small last chunk": {input: []T{8, 7, 5, 3, 3, -15, 10}, fold: utils.Add[T], expects: 21},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := palgo.Reduce(tc.input, tc.fold, 0)
			require.Equal(t, tc.expects, got)
		})
	}
}

func testMapReduce[T utils.Signed](t *testing.T) { //nolint: thelper
	t.Parallel()

	sq := func(x T) int { return int(x * x) }
	acc := func(x T, y int) T { return x + T(y) }

	testCases := map[string]struct {
		input   []T
		unary   func(T) int
		fold    func(T, int) T
		expects T
	}{
		"empty sum of squares":  {input: []T{}, unary: sq, fold: acc, expects: 0},
		"small sum of squares":  {input: []T{1, 2, 3}, unary: sq, fold: acc, expects: 14},
		"normal sum of squares": {input: []T{1, 5, 3, 3, -2, 5, 1, 2, 4}, unary: sq, fold: acc, expects: 94},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := algo.MapReduce(tc.input, tc.unary, tc.fold, 0)
			require.Equal(t, tc.expects, got)
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
			got := algo.ZipWith(tc.input1, tc.input2, tc.zip)
			require.Equal(t, tc.want, got)
		})
	}
}

func testZipReduce[T utils.Signed](t *testing.T) { //nolint: thelper
	t.Parallel()

	testCases := map[string]struct {
		input1 []T
		input2 []T
		want   T
	}{
		"empty":      {want: 0},
		"half empty": {input1: []T{1}, want: 0},
		"single":     {input1: []T{1}, input2: []T{2}, want: 2},
		"normal":     {input1: []T{1, 3, 9}, input2: []T{2, -1, 6}, want: 53},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			// Testing by computing inner product.
			got := algo.ZipReduce(tc.input1, tc.input2, utils.Mul[T], utils.Add[T], 0)
			require.Equal(t, tc.want, got)
		})
	}
}
