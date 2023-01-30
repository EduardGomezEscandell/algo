package dstruct

type lruEntry[K, V any] struct {
	epoch epoch
	key   K
	data  V
}

type index int
type epoch uint64

// LruCache is a implements a Least Recently Used cache.
// It acts as a dictionary of keys and values, with a maximum number of
// entries. When this maximum is surpassed, the least recently used item
// is dropped.
type LruCache[K comparable, V any] struct {
	byAge    Heap[index]      // Heap to quickly acces items by age.
	byKey    map[K]index      // Map to quickly access to items by key.
	data     []lruEntry[K, V] // Raw data.
	capacity int              // Max amount of items.
	epoch    epoch            // A timestamp. Updated every read and write.
}

// NewLRU creates new lru cache with the specified capacity,
// measured in number of key-value paires stored.
func NewLRU[K comparable, V any](cap int) *LruCache[K, V] {
	alloc := make([]lruEntry[K, V], cap)

	return &LruCache[K, V]{
		byKey:    map[K]index{},
		byAge:    NewHeap(func(x, y index) bool { return alloc[x].epoch < alloc[y].epoch }),
		data:     alloc,
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
	entry, ok := lru.get(key)
	if !ok {
		return v, false
	}
	return entry.data, true
}

// get accesses an item in the cache and updates its epoch.
func (lru *LruCache[K, V]) get(k K) (*lruEntry[K, V], bool) {
	idx, ok := lru.byKey[k]
	entry := &lru.data[idx]
	if !ok {
		return nil, false
	}
	lru.epoch++
	entry.epoch = lru.epoch
	lru.byAge.Fix(int(idx))

	return entry, true
}

// Set checks if an entry was in the cache. If it was,
// it updates its epoch. Otherwise, it adds it anew.
func (lru LruCache[K, V]) Set(k K, v V) {
	lru.epoch++
	entry, ok := lru.get(k)
	if ok {
		entry.data = v
		return
	}

	var idx index
	if lru.Len() >= lru.capacity {
		idx = lru.byAge.Pop()
		delete(lru.byKey, lru.data[idx].key)
	} else {
		idx = index(lru.Len())
	}
	ptr := &lru.data[idx]

	*ptr = lruEntry[K, V]{
		epoch: lru.epoch,
		key:   k,
		data:  v,
	}

	lru.byKey[k] = idx
	lru.byAge.Push(idx)
}
