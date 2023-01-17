package algo_test

import (
	"testing"

	"github.com/EduardGomezEscandell/algo/algo"
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

func TestAdjacentMap(t *testing.T) {
	t.Parallel()
	t.Run("int", testAdjacentMap[int])
	t.Run("int8", testAdjacentMap[int8])
	t.Run("int32", testAdjacentMap[int32])
	t.Run("int64", testAdjacentMap[int64])
}

func TestAdjacentReduce(t *testing.T) {
	t.Parallel()
	t.Run("int", testAdjacentReduce[int])
	t.Run("int8", testAdjacentReduce[int8])
	t.Run("int32", testAdjacentReduce[int32])
	t.Run("int64", testAdjacentReduce[int64])
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

func TestTopN(t *testing.T) {
	t.Parallel()
	t.Run("int", testTopN[int])
	t.Run("int8", testTopN[int8])
	t.Run("int32", testTopN[int32])
	t.Run("int64", testTopN[int64])
}

func TestCommon(t *testing.T) {
	t.Parallel()
	t.Run("int", testCommon[int])
	t.Run("int8", testCommon[int8])
	t.Run("int32", testCommon[int32])
	t.Run("int64", testCommon[int64])
}

func TestUnique(t *testing.T) {
	t.Parallel()
	t.Run("int", testUnique[int])
	t.Run("int8", testUnique[int8])
	t.Run("int32", testUnique[int32])
	t.Run("int64", testUnique[int64])
}

func TestFind(t *testing.T) {
	t.Parallel()
	t.Run("int", testFind[int])
	t.Run("int8", testFind[int8])
	t.Run("int32", testFind[int32])
	t.Run("int64", testFind[int64])
}

func TestPartition(t *testing.T) {
	t.Parallel()
	t.Run("int", testPartition[int])
	t.Run("int8", testPartition[int8])
	t.Run("int32", testPartition[int32])
	t.Run("int64", testPartition[int64])
}

func TestInsert(t *testing.T) {
	t.Parallel()
	t.Run("int", testInsert[int])
	t.Run("int8", testInsert[int8])
	t.Run("int32", testInsert[int32])
	t.Run("int64", testInsert[int64])
}

func TestRotate(t *testing.T) {
	t.Parallel()
	t.Run("int", testRotate[int])
	t.Run("int8", testRotate[int8])
	t.Run("int32", testRotate[int32])
	t.Run("int64", testRotate[int64])
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
			got := algo.Map(tc.input, tc.op)
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
		"empty sum":          {input: []T{}, fold: utils.Add[T], expects: 0},
		"small sum":          {input: []T{1, 2, 3}, fold: utils.Add[T], expects: 6},
		"normal sum":         {input: []T{8, 7, 5, 3, 3, -15}, fold: utils.Add[T], expects: 11},
		"empty subtraction":  {input: []T{}, fold: utils.Sub[T], expects: 0},
		"small subtraction":  {input: []T{1, 2, 3}, fold: utils.Sub[T], expects: -6},
		"normal subtraction": {input: []T{8, 7, 5, 3, 3, -15}, fold: utils.Sub[T], expects: -11},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := algo.Reduce(tc.input, tc.fold, 0)
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
		"normal sum of squares": {input: []T{7, 5, 3, 3}, unary: sq, fold: acc, expects: 92},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := algo.MapReduce(tc.input, tc.unary, tc.fold, 0)
			require.Equal(t, tc.expects, got)
		})
	}
}

func testAdjacentReduce[T utils.Signed](t *testing.T) { //nolint: thelper
	t.Parallel()
	testCases := map[string]struct {
		input   []T
		merge   func(T, T) T
		fold    func(T, T) T
		expects T
	}{
		"empty sum of differences":  {merge: utils.Sub[T], fold: utils.Add[T], expects: 0, input: []T{}},
		"small sum of differences":  {merge: utils.Sub[T], fold: utils.Add[T], expects: -2, input: []T{1, 2, 3}},
		"normal sum of differences": {merge: utils.Sub[T], fold: utils.Add[T], expects: 23, input: []T{8, 7, 5, 3, 3, -15}},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := algo.AdjacentReduce(tc.input, tc.merge, tc.fold)
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

func testAdjacentMap[T utils.Signed](t *testing.T) { //nolint: thelper
	t.Parallel()
	testCases := map[string]struct {
		input   []T
		op      func(T, T) T
		expects []T
	}{
		"empty sum":          {input: []T{}, op: utils.Add[T], expects: []T{}},
		"small sum":          {input: []T{1, -2, 3}, op: utils.Add[T], expects: []T{-1, 1}},
		"normal sum":         {input: []T{-8, 7, 0, 3, 3, -15}, op: utils.Add[T], expects: []T{-1, 7, 3, 6, -12}},
		"empty subtraction":  {input: []T{}, op: utils.Sub[T], expects: []T{}},
		"small subtraction":  {input: []T{1, -2, 3}, op: utils.Sub[T], expects: []T{3, -5}},
		"normal subtraction": {input: []T{-8, 7, 0, 3, 3, -15}, op: utils.Sub[T], expects: []T{-15, 7, -3, 0, 18}},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := algo.AdjacentMap(tc.input, tc.op)
			require.Equal(t, tc.expects, got)
		})
	}
}

func testTopN[T utils.Signed](t *testing.T) { //nolint: thelper
	t.Parallel()
	testCases := map[string]struct {
		input []T
		n     uint
		comp  utils.Comparator[T]
		want  []T
	}{
		"empty":           {n: 2, input: []T{}, comp: utils.Lt[T], want: []T{}},
		"too short":       {n: 5, input: []T{1, 2, 3}, comp: utils.Lt[T], want: []T{1, 2, 3}},
		"just the amount": {n: 5, input: []T{1, 4, 2, 3}, comp: utils.Lt[T], want: []T{1, 2, 3, 4}},
		"bottom 3":        {n: 3, input: []T{8, 7, 5, 3, 3, 15}, comp: utils.Lt[T], want: []T{3, 3, 5}},
		"top 3":           {n: 3, input: []T{8, 7, 5, 3, 3, 15}, comp: utils.Gt[T], want: []T{15, 8, 7}},
		"bottom 4":        {n: 4, input: []T{-1, 84, 5, 101, 12, 9, 15, 1}, comp: utils.Lt[T], want: []T{-1, 1, 5, 9}},
		"top 4":           {n: 4, input: []T{-1, 84, 5, 101, 12, 9, 15, 1}, comp: utils.Gt[T], want: []T{101, 84, 15, 12}},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := algo.FirstN(tc.input, tc.n, tc.comp)
			require.Equal(t, tc.want, got)
		})
	}
}

func testCommon[T utils.Signed](t *testing.T) { //nolint: thelper
	t.Parallel()
	testCases := map[string]struct {
		input1 []T
		input2 []T
		sort   func(T, T) bool
		want   []T
	}{
		"less than, empty":                               {sort: utils.Lt[T], want: []T{}},
		"greater than, empty":                            {sort: utils.Gt[T], want: []T{}},
		"less than, half empty":                          {sort: utils.Lt[T], input1: []T{1, 6, 8}, want: []T{}},
		"greater than, half empty":                       {sort: utils.Gt[T], input1: []T{1, 6, 8}, want: []T{}},
		"less than, single, no shared":                   {sort: utils.Lt[T], input1: []T{1}, input2: []T{2}, want: []T{}},
		"greater than, single, no shared":                {sort: utils.Gt[T], input1: []T{1}, input2: []T{2}, want: []T{}},
		"less than, single, shared":                      {sort: utils.Lt[T], input1: []T{1}, input2: []T{1}, want: []T{1}},
		"greater than, single, shared":                   {sort: utils.Gt[T], input1: []T{1}, input2: []T{1}, want: []T{1}},
		"less than, normal, shared":                      {sort: utils.Lt[T], input1: []T{1, 3, 9}, input2: []T{1, 5, 9}, want: []T{1, 9}},
		"greater than, normal, shared":                   {sort: utils.Gt[T], input1: []T{1, 3, 9}, input2: []T{1, 5, 9}, want: []T{9, 1}},
		"less than, normal, no shared":                   {sort: utils.Lt[T], input1: []T{1, 3, 9}, input2: []T{2, 4, 6}, want: []T{}},
		"greater than, normal, no shared":                {sort: utils.Gt[T], input1: []T{1, 3, 9}, input2: []T{2, 4, 6}, want: []T{}},
		"less than, repeats, no shared":                  {sort: utils.Lt[T], input1: []T{15, 0, 15, 9}, input2: []T{0, 15}, want: []T{0, 15}},
		"greater than, repeats, no shared":               {sort: utils.Gt[T], input1: []T{15, 0, 15, 9}, input2: []T{0, 15}, want: []T{15, 0}},
		"less than, double repeats, no shared":           {sort: utils.Lt[T], input1: []T{15, 0, 15, 9}, input2: []T{0, 15, 15}, want: []T{0, 15, 15}},
		"greater than, double repeats, no shared":        {sort: utils.Gt[T], input1: []T{15, 0, 15, 9}, input2: []T{0, 15, 15}, want: []T{15, 15, 0}},
		"greater than, triple double repeats, no shared": {sort: utils.Gt[T], input1: []T{15, 15, 15, 9}, input2: []T{0, 15, 15}, want: []T{15, 15}},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			algo.Sort(tc.input1, tc.sort)
			algo.Sort(tc.input2, tc.sort)
			got := algo.Intersect(tc.input1, tc.input2, tc.sort)
			require.Equal(t, tc.want, got)
		})
	}
}

func testUnique[T utils.Signed](t *testing.T) { //nolint: thelper
	t.Parallel()
	testCases := map[string]struct {
		input     []T
		comp      func(T, T) bool
		wantArr   []T
		wantCount int
	}{
		"empty":                           {comp: utils.Lt[T]},
		"one":                             {comp: utils.Lt[T], input: []T{5}, wantArr: []T{5}, wantCount: 1},
		"few, greater than, no repeats":   {comp: utils.Gt[T], input: []T{3, 1, 5}, wantArr: []T{5, 3, 1}, wantCount: 3},
		"few, less than, no repeats":      {comp: utils.Lt[T], input: []T{3, 1, 5}, wantArr: []T{1, 3, 5}, wantCount: 3},
		"few, greater than, some repeats": {comp: utils.Gt[T], input: []T{3, 1, 3, 5}, wantArr: []T{5, 3, 1, 3}, wantCount: 3},
		"few, less than, some repeats":    {comp: utils.Lt[T], input: []T{3, 1, 3, 5}, wantArr: []T{1, 3, 5, 3}, wantCount: 3},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			algo.Sort(tc.input, tc.comp)
			got := algo.Unique(tc.input, utils.Equal(tc.comp))
			require.Equal(t, tc.wantArr, tc.input)
			require.Equal(t, tc.wantCount, got)
		})
	}
}

func testFind[T utils.Number](t *testing.T) { //nolint: thelper
	t.Parallel()
	testCases := map[string]struct {
		data   []T
		search T
		want   int
	}{
		"empty":          {data: []T{}, want: -1},
		"single, found":  {data: []T{5}, search: 5, want: 0},
		"single, missed": {data: []T{3}, search: 9, want: -1},
		"missed":         {data: []T{3, 13, 84, 6, 3}, search: 9, want: -1},
		"first":          {data: []T{3, 13, 84, 6, 11}, search: 3, want: 0},
		"last":           {data: []T{3, 13, 84, 6, 11}, search: 11, want: 4},
		"repeated":       {data: []T{3, 13, 84, 6, 3}, search: 3, want: 0},
		"standard":       {data: []T{3, 13, 84, 6, 3}, search: 84, want: 2},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			input := make([]T, len(tc.data))
			copy(input, tc.data)

			got := algo.Find(input, tc.search, utils.Eq[T])
			require.Equal(t, tc.want, got)

			require.Equal(t, input, tc.data, "algo.Find modified the input array")
		})
	}
}

func testPartition[T utils.Number](t *testing.T) { //nolint: thelper
	t.Parallel()
	testCases := map[string]struct {
		data   []T
		search T
		want   int
	}{
		"empty":          {data: []T{}, want: 0},
		"single, before": {data: []T{5}, search: 3, want: 0},
		"single, after":  {data: []T{3}, search: 9, want: 1},
		"standard":       {data: []T{3, 13, 84, 6, 3}, search: 9, want: 3},
		"bug":            {data: []T{1, 2, 15}, search: 9, want: 2},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			input := make([]T, len(tc.data))
			copy(input, tc.data)

			got := algo.Partition(input, func(x T) bool { return x < tc.search })
			require.Equal(t, tc.want, got)

			for i, v := range input[:got] {
				require.LessOrEqual(t, v, tc.search, "Item %d is greater than the partition: %d > %d. Array: %v", i, v, tc.search, input)
			}
			for i, v := range input[got:] {
				require.GreaterOrEqual(t, v, tc.search, "Item %d is smaller than the partition: %d < %d. Array: %v", i, v, tc.search, input)
			}
		})
	}
}

func testInsert[T utils.Signed](t *testing.T) { //nolint: thelper
	t.Parallel()

	testCases := map[string]struct {
		input []T
		pos   int
		val   T
		want  []T
	}{
		"empty":  {input: []T{}, val: 5, pos: 0, want: []T{5}},
		"lead":   {input: []T{12, 8, 6}, val: 11, pos: 0, want: []T{11, 12, 8, 6}},
		"tail":   {input: []T{12, 8, 6}, val: 11, pos: 3, want: []T{12, 8, 6, 11}},
		"middle": {input: []T{12, 8, 6}, val: 11, pos: 1, want: []T{12, 11, 8, 6}},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			in := make([]T, len(tc.input))
			copy(in, tc.input)
			got := algo.Insert(in, tc.val, tc.pos)
			require.Equal(t, tc.want, got)
		})
	}
}

func testRotate[T utils.Signed](t *testing.T) { //nolint: thelper
	t.Parallel()

	testCases := map[string]struct {
		input   []T
		n       int
		wantArr []T
		wantIdx int
	}{
		// "single": {input: []T{5}, n: 0, wantArr: []T{5}, wantIdx: 0},
		// "left": {input: []T{1, 3, 5, 7, 9}, n: 2, wantArr: []T{5, 7, 9, 1, 3}, wantIdx: 3},
		"right": {input: []T{1, 3, 5, 7, 9}, n: -2, wantArr: []T{7, 9, 1, 3, 5}, wantIdx: 2},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			arr := make([]T, len(tc.input))
			copy(arr, tc.input)
			got := algo.Rotate(arr, tc.n)
			require.Equal(t, tc.wantArr, arr)
			require.Equal(t, tc.wantIdx, got)
		})
	}
}
