package dstruct

// LruCache is a implements a Least Recently Used cache.
// It acts as a dictionary of keys and values, with a maximum number of
// entries. When this maximum is surpassed, the least recently used item
// is dropped.
type LruCache[K comparable, V any] struct {
	byAge    Heap[*lruEntry[K, V]] // Heap to quickly acces items by age.
	byKey    map[K]*lruEntry[K, V] // Map to quickly access to items by key.
	data     []lruEntry[K, V]      // Raw data.
	capacity int                   // Max amount of items.
	epoch    lruEpoch              // A timestamp. Updated every read and write.
}

// NewLRU creates new lru cache with the specified capacity,
// measured in number of key-value paires stored.
func NewLRU[K comparable, V any](cap int) *LruCache[K, V] {
	return &LruCache[K, V]{
		byKey:    map[K]*lruEntry[K, V]{},
		byAge:    NewHeap(func(x, y *lruEntry[K, V]) bool { return x.epoch < y.epoch }),
		data:     make([]lruEntry[K, V], cap),
		capacity: cap,
		epoch:    1,
	}
}

// Len is the number of stored items.
func (lru LruCache[K, V]) Len() int {
	return len(lru.byKey)
}

// Get looks into the cache to see if the given key is
// is registered. If so, the value is returned and its
// epoch updated.
func (lru *LruCache[K, V]) Get(key K) (v V, ok bool) {
	entry, ok := lru.byKey[key]
	if !ok {
		return v, false
	}
	lru.epoch++
	entry.epoch = lru.epoch
	return entry.data, true
}

// Set checks if an entry was in the cache. If it was,
// it updates its epoch. Otherwise, it adds it anew.
func (lru LruCache[K, V]) Set(k K, v V) {
	lru.epoch++
	entry, ok := lru.byKey[k]
	if ok {
		entry.epoch = lru.epoch
		return
	}
	var ptr *lruEntry[K, V]
	if lru.Len() >= lru.capacity {
		ptr = lru.byAge.Pop()
		delete(lru.byKey, ptr.key)
	} else {
		ptr = &lru.data[lru.Len()]
	}

	ptr.epoch = lru.epoch
	ptr.key = k
	ptr.data = v

	lru.byKey[k] = ptr
	lru.byAge.Push(ptr)
}

type lruEpoch uint64

type lruEntry[K, V any] struct {
	epoch lruEpoch
	key   K
	data  V
}
