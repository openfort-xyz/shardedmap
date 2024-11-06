package shardedmap

import (
	"testing"
)

func TestShardedMap_SetAndGet(t *testing.T) {
	m := NewShardedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")

	if val, ok := m.Get(1); !ok || val != "one" {
		t.Errorf("expected 'one', got '%v'", val)
	}

	if val, ok := m.Get(2); !ok || val != "two" {
		t.Errorf("expected 'two', got '%v'", val)
	}

	if _, ok := m.Get(3); ok {
		t.Errorf("expected key 3 to be absent")
	}
}

func TestShardedMap_Delete(t *testing.T) {
	m := NewShardedMap[int, string]()
	m.Set(1, "one")
	m.Delete(1)

	if _, ok := m.Get(1); ok {
		t.Errorf("expected key 1 to be deleted")
	}
}

func TestShardedMap_Count(t *testing.T) {
	m := NewShardedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")

	if count := m.Count(); count != 2 {
		t.Errorf("expected count to be 2, got %d", count)
	}

	m.Delete(1)

	if count := m.Count(); count != 1 {
		t.Errorf("expected count to be 1, got %d", count)
	}
}

func TestShardedMap_Keys(t *testing.T) {
	m := NewShardedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")

	keys := m.Keys()
	expectedKeys := map[int]bool{1: true, 2: true}

	for _, key := range keys {
		if !expectedKeys[key] {
			t.Errorf("unexpected key %d", key)
		}
		delete(expectedKeys, key)
	}

	if len(expectedKeys) != 0 {
		t.Errorf("expected keys %v not found", expectedKeys)
	}
}

func TestShardedMap_Clear(t *testing.T) {
	m := NewShardedMap[int, string]()
	m.Set(1, "one")
	m.Set(2, "two")
	m.Clear()

	if count := m.Count(); count != 0 {
		t.Errorf("expected count to be 0 after clear, got %d", count)
	}

	if keys := m.Keys(); len(keys) != 0 {
		t.Errorf("expected no keys after clear, got %v", keys)
	}
}

func TestWithShardCount(t *testing.T) {
	shardCount := uint32(64)
	m := NewShardedMap[int, string](WithShardCount(shardCount))

	if len(m.shards) != int(shardCount) {
		t.Errorf("expected shard count to be %d, got %d", shardCount, len(m.shards))
	}
}
