package shardedmap

type options struct {
	ShardCount *uint32
}

// Option is a functional option for configuring a sharded map.
type Option func(*options)

// WithShardCount sets the number of shards in the sharded map.
func WithShardCount(shardCount uint32) Option {
	return func(o *options) {
		o.ShardCount = &shardCount
	}
}
