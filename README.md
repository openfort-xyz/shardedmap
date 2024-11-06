# shardedmap
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)


```go
import "go.openfort.xyz/shardedmap"
```

A Go package that implements a thread-safe, sharded map with configurable concurrency for optimized key-value storage operations.

## Index

- [Variables](<#variables>)
- [type MapShard](<#MapShard>)
- [type Option](<#Option>)
  - [func WithShardCount\(shardCount int\) Option](<#WithShardCount>)
- [type ShardedMap](<#ShardedMap>)
  - [func NewShardedMap\[K comparable, V any\]\(opts ...Option\) \*ShardedMap\[K, V\]](<#NewShardedMap>)
  - [func \(m \*ShardedMap\[K, V\]\) Clear\(\)](<#ShardedMap[K, V].Clear>)
  - [func \(m \*ShardedMap\[K, V\]\) Count\(\) int32](<#ShardedMap[K, V].Count>)
  - [func \(m \*ShardedMap\[K, V\]\) Delete\(key K\)](<#ShardedMap[K, V].Delete>)
  - [func \(m \*ShardedMap\[K, V\]\) Get\(key K\) \(V, bool\)](<#ShardedMap[K, V].Get>)
  - [func \(m \*ShardedMap\[K, V\]\) Keys\(\) \[\]K](<#ShardedMap[K, V].Keys>)
  - [func \(m \*ShardedMap\[K, V\]\) Set\(key K, value V\)](<#ShardedMap[K, V].Set>)


## Variables

<a name="ShardCount"></a>ShardCount is the default number of shards used by the ShardedMap. This can be adjusted based on your concurrency requirements using the WithShardCount option.

```go
var ShardCount = 32
```

<a name="MapShard"></a>
## type MapShard

MapShard represents a single shard in the ShardedMap. It contains a map and a read\-write mutex to ensure thread\-safe operations.

```go
type MapShard[K comparable, V any] struct {
    // contains filtered or unexported fields
}
```

<a name="Option"></a>
## type Option

Option is a functional option for configuring a sharded map.

```go
type Option func(*options)
```

<a name="WithShardCount"></a>
### func WithShardCount

```go
func WithShardCount(shardCount int) Option
```

WithShardCount sets the number of shards in the sharded map.

<a name="ShardedMap"></a>
## type ShardedMap

ShardedMap is a thread\-safe map where keys are distributed across multiple shards.

```go
type ShardedMap[K comparable, V any] struct {
    // contains filtered or unexported fields
}
```

<a name="NewShardedMap"></a>
### func NewShardedMap

```go
func NewShardedMap[K comparable, V any](opts ...Option) *ShardedMap[K, V]
```

NewShardedMap initializes a ShardedMap with the specified options or defaults.

<a name="ShardedMap[K, V].Clear"></a>
### func \(\*ShardedMap\[K, V\]\) Clear

```go
func (m *ShardedMap[K, V]) Clear()
```

Clear removes all key\-value pairs from the map.

<a name="ShardedMap[K, V].Count"></a>
### func \(\*ShardedMap\[K, V\]\) Count

```go
func (m *ShardedMap[K, V]) Count() int32
```

Count returns the number of elements in the map across all shards.

<a name="ShardedMap[K, V].Delete"></a>
### func \(\*ShardedMap\[K, V\]\) Delete

```go
func (m *ShardedMap[K, V]) Delete(key K)
```

Delete removes a key from the map.

<a name="ShardedMap[K, V].Get"></a>
### func \(\*ShardedMap\[K, V\]\) Get

```go
func (m *ShardedMap[K, V]) Get(key K) (V, bool)
```

Get retrieves a value from the map by key. The second return value indicates whether the key was found.

<a name="ShardedMap[K, V].Keys"></a>
### func \(\*ShardedMap\[K, V\]\) Keys

```go
func (m *ShardedMap[K, V]) Keys() []K
```

Keys returns a slice containing all the keys in the map. This method is thread\-safe, but may be relatively slow depending on the map size.

<a name="ShardedMap[K, V].Set"></a>
### func \(\*ShardedMap\[K, V\]\) Set

```go
func (m *ShardedMap[K, V]) Set(key K, value V)
```

Set sets a key\-value pair in the map. It overwrites any existing value for the key.
