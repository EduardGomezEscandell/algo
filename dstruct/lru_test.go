package dstruct_test

import (
	"testing"

	"github.com/EduardGomezEscandell/algo/dstruct"
	"github.com/stretchr/testify/require"
)

func TestLRU(t *testing.T) {
	lru := dstruct.NewLRU[uint, string](5)
	lru.Set(1, "Numero uno")

	v, ok := lru.Get(1)
	require.True(t, ok, "Failed to get value that should have been in cache")
	require.Equal(t, v, "Numero uno", "Retrieved wrong value from cache")

	lru.Set(1, "one")
	lru.Set(2, "two")
	lru.Set(3, "three")

	v, ok = lru.Get(1)
	require.True(t, ok, "Failed to get value that should have been in cache")
	require.Equal(t, v, "one", "Retrieved wrong value from cache")

	v, ok = lru.Get(3)
	require.True(t, ok, "Failed to get value that should have been in cache")
	require.Equal(t, v, "three", "Retrieved wrong value from cache")

	_, ok = lru.Get(700)
	require.False(t, ok, "Got value that should have not been in cache")

	v, ok = lru.Get(2)
	require.True(t, ok, "Failed to get value that should have been in cache")
	require.Equal(t, v, "two", "Retrieved wrong value from cache")

	// We ensure that {2, "two"} and {3, "three"} are removed from cache when newer entries are added.
	// We also ensure that the newer entries are properly written in.
	lru.Get(1) // Refreshed
	lru.Set(100, "one hundred")
	lru.Set(200, "two hundred")
	lru.Set(300, "three hundred")
	lru.Set(400, "four hundred")

	v, ok = lru.Get(1)
	require.True(t, ok, "Failed to get value that should have been in cache")
	require.Equal(t, v, "one", "Retrieved wrong value from cache")

	_, ok = lru.Get(2)
	require.False(t, ok, "Got value that have been removed from cache")

	_, ok = lru.Get(3)
	require.False(t, ok, "Got value that have been removed from cache")

	v, ok = lru.Get(100)
	require.True(t, ok, "Failed to get value that should have been in cache")
	require.Equal(t, v, "one hundred", "Retrieved wrong value from cache")

	v, ok = lru.Get(200)
	require.True(t, ok, "Failed to get value that should have been in cache")
	require.Equal(t, v, "two hundred", "Retrieved wrong value from cache")

	v, ok = lru.Get(300)
	require.True(t, ok, "Failed to get value that should have been in cache")
	require.Equal(t, v, "three hundred", "Retrieved wrong value from cache")

	v, ok = lru.Get(400)
	require.True(t, ok, "Failed to get value that should have been in cache")
	require.Equal(t, v, "four hundred", "Retrieved wrong value from cache")

	// We ensure that {1, "one"} has been refreshed when we last accessed it.
	lru.Get(1) // Refreshed
	lru.Set(1000, "one thousand")
	lru.Set(2000, "two thousand")
	lru.Set(3000, "three thousand")
	lru.Set(4000, "four thousand")

	v, ok = lru.Get(1)
	require.True(t, ok, "Failed to get value that should have been in cache")
	require.Equal(t, v, "one", "Retrieved wrong value from cache")
}
