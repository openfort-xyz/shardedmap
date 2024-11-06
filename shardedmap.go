package shardedmap

import (
	"fmt"
	"hash/fnv"
	"sync"
	"sync/atomic"
)

// ShardCount is the default number of shards used by the ShardedMap.
// This can be adjusted based on your concurrency requirements using the WithShardCount option.
var ShardCount uint32 = 32

// MapShard represents a single shard in the ShardedMap.
// It contains a map and a read-write mutex to ensure thread-safe operations.
type MapShard[K comparable, V any] struct {
	mu    sync.RWMutex
	data  map[K]V
	count int32
}

// ShardedMap is a thread-safe map where keys are distributed across multiple shards.
type ShardedMap[K comparable, V any] struct {
	shards []*MapShard[K, V]
}

// NewShardedMap initializes a ShardedMap with the specified options or defaults.
func NewShardedMap[K comparable, V any](opts ...Option) *ShardedMap[K, V] {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}

	if o.ShardCount != nil && *o.ShardCount > 0 {
		ShardCount = *o.ShardCount
	}

	shards := make([]*MapShard[K, V], ShardCount)
	for i := uint32(0); i < ShardCount; i++ {
		shards[i] = &MapShard[K, V]{
			data: make(map[K]V),
		}
	}
	return &ShardedMap[K, V]{shards: shards}
}

// getShard returns the shard corresponding to a given key.
func (m *ShardedMap[K, V]) getShard(key K) *MapShard[K, V] {
	h := fnv.New32()
	h.Write([]byte(fmt.Sprintf("%v", key))) // Using fmt.Sprintf to convert key to string
	return m.shards[h.Sum32()%ShardCount]
}

// Set sets a key-value pair in the map. It overwrites any existing value for the key.
func (m *ShardedMap[K, V]) Set(key K, value V) {
	shard := m.getShard(key)
	shard.mu.Lock()
	defer shard.mu.Unlock()
	if _, exists := shard.data[key]; !exists {
		atomic.AddInt32(&shard.count, 1)
	}
	shard.data[key] = value
}

// Get retrieves a value from the map by key. The second return value indicates whether the key was found.
func (m *ShardedMap[K, V]) Get(key K) (V, bool) {
	shard := m.getShard(key)
	shard.mu.RLock()
	defer shard.mu.RUnlock()
	val, ok := shard.data[key]
	return val, ok
}

// Delete removes a key from the map.
func (m *ShardedMap[K, V]) Delete(key K) {
	shard := m.getShard(key)
	shard.mu.Lock()
	defer shard.mu.Unlock()
	if _, exists := shard.data[key]; exists {
		atomic.AddInt32(&shard.count, -1)
		delete(shard.data, key)
	}
}

// Count returns the number of elements in the map across all shards.
func (m *ShardedMap[K, V]) Count() int32 {
	var total int32
	for _, shard := range m.shards {
		total += atomic.LoadInt32(&shard.count)
	}
	return total
}

// Keys returns a slice containing all the keys in the map.
// This method is thread-safe, but may be relatively slow depending on the map size.
func (m *ShardedMap[K, V]) Keys() []K {
	keys := make([]K, 0)
	for _, shard := range m.shards {
		shard.mu.RLock()
		for key := range shard.data {
			keys = append(keys, key)
		}
		shard.mu.RUnlock()
	}
	return keys
}

// Clear removes all key-value pairs from the map.
func (m *ShardedMap[K, V]) Clear() {
	for _, shard := range m.shards {
		shard.mu.Lock()
		shard.data = make(map[K]V)
		atomic.StoreInt32(&shard.count, 0)
		shard.mu.Unlock()
	}
}
