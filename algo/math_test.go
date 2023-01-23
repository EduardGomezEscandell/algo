package algo_test

import (
	"fmt"
	"testing"

	"github.com/EduardGomezEscandell/algo/algo"
	"github.com/EduardGomezEscandell/algo/utils"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

func TestAbs(t *testing.T) {
	t.Parallel()
	t.Run("int", testAbs[int])
	t.Run("int8", testAbs[int8])
	t.Run("int32", testAbs[int32])
	t.Run("int64", testAbs[int64])
}

func TestGCD(t *testing.T) {
	t.Parallel()
	t.Run("int", testGCD[int])
	t.Run("int8", testGCD[int8])
	t.Run("int32", testGCD[int32])
	t.Run("int64", testGCD[int64])

	t.Run("uint", testGCD[uint])
	t.Run("uint8", testGCD[uint8])
	t.Run("uint32", testGCD[uint32])
	t.Run("uint64", testGCD[uint64])
}

func TestLCM(t *testing.T) {
	t.Parallel()
	t.Run("int", testLCM[int])
	t.Run("int8", testLCM[int8])
	t.Run("int32", testLCM[int32])
	t.Run("int64", testLCM[int64])

	t.Run("uint", testLCM[uint])
	t.Run("uint8", testLCM[uint8])
	t.Run("uint32", testLCM[uint32])
	t.Run("uint64", testLCM[uint64])
}

func TestClamp(t *testing.T) {
	t.Parallel()
	t.Run("int", testClamp[int])
	t.Run("int8", testClamp[int8])
	t.Run("int32", testClamp[int32])
	t.Run("int64", testClamp[int64])

	t.Run("uint", testClamp[uint])
	t.Run("uint8", testClamp[uint8])
	t.Run("uint32", testClamp[uint32])
	t.Run("uint64", testClamp[uint64])
}

func testAbs[T utils.Signed](t *testing.T) { //nolint: thelper
	t.Parallel()
	testCases := map[T]T{
		0:   0,
		13:  13,
		-15: 15,
	}

	for input, want := range testCases {
		input, want := input, want
		t.Run(fmt.Sprintf("%v", input), func(t *testing.T) {
			require.Equal(t, want, algo.Abs(input))
		})
	}
}

func testGCD[T constraints.Integer](t *testing.T) { //nolint: thelper
	t.Parallel()
	testCases := map[string]struct {
		input [2]T
		want  T
	}{
		"couple":    {input: [2]T{15, 25}, want: 5},
		"coprime":   {input: [2]T{2, 5}, want: 1},
		"contained": {input: [2]T{7, 14}, want: 7},
		"nonprime":  {input: [2]T{99, 27}, want: 9},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := algo.GCD(tc.input[0], tc.input[1])
			require.Equal(t, tc.want, got, "unexpected result of GCD")

			got = algo.GCD(tc.input[1], tc.input[0])
			require.Equal(t, tc.want, got, "GCD failed commutativity test")
		})
	}
}

func testLCM[T constraints.Integer](t *testing.T) { //nolint: thelper
	t.Parallel()
	testCases := map[string]struct {
		input []T
		want  T
	}{
		"couple":     {input: []T{9, 6}, want: 18},
		"coprime":    {input: []T{9, 7}, want: 63},
		"contained":  {input: []T{7, 14}, want: 14},
		"triplet":    {input: []T{2, 3, 4}, want: 12},
		"quadruplet": {input: []T{2, 3, 4, 5}, want: 60},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := algo.LCM(tc.input[0], tc.input[1], tc.input[2:]...)
			require.Equal(t, tc.want, got, "unexpected result of LCM")

			got = algo.LCM(tc.input[1], tc.input[0], tc.input[2:]...)
			require.Equal(t, tc.want, got, "LCM failed commutativity test")
		})
	}
}

func testClamp[T constraints.Integer](t *testing.T) { //nolint: thelper
	t.Parallel()
	testCases := map[string]struct {
		input  T
		limits [2]T
		want   T
	}{
		"contained":   {input: 3, limits: [2]T{0, 5}, want: 3},
		"lower bound": {input: 12, limits: [2]T{12, 99}, want: 12},
		"upper bound": {input: 55, limits: [2]T{16, 55}, want: 55},
		"below":       {input: 13, limits: [2]T{64, 99}, want: 64},
		"above":       {input: 61, limits: [2]T{12, 15}, want: 15},
		"closed":      {input: 6, limits: [2]T{5, 5}, want: 5},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := algo.Clamp(tc.limits[0], tc.input, tc.limits[1])
			require.Equal(t, tc.want, got, "unexpected result of Clamp")
		})
	}
}

func TestCount(t *testing.T) {
	t.Parallel()
	require.Equal(t, 2, algo.Count(1, true))
	require.Equal(t, 1, algo.Count(1, false))
}
