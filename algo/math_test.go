package algo_test

import (
	"fmt"
	"testing"

	"github.com/EduardGomezEscandell/algo/algo"
	"github.com/EduardGomezEscandell/algo/utils"
	"github.com/stretchr/testify/require"
)

func TestAbs(t *testing.T) {
	t.Parallel()
	t.Run("int", testAbs[int])
	t.Run("int8", testAbs[int8])
	t.Run("int32", testAbs[int32])
	t.Run("int64", testAbs[int64])
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
