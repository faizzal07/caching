package cache

import (
	"testing"
	"time"
)

func TestLRUCache_SetAndGet(t *testing.T) {
    cache := NewLRUCache(2)

    cache.Set("key1", "value1", 1*time.Minute)
    cache.Set("key2", "value2", 1*time.Minute)

    if value, ok := cache.Get("key1"); !ok || value != "value1" {
        t.Errorf("expected value1, got %v", value)
    }

    cache.Set("key3", "value3", 1*time.Minute) // This should evict key2 due to LRU policy

    if _, ok := cache.Get("key2"); ok {
        t.Error("expected key2 to be evicted")
    }
}

func TestLRUCache_TTL(t *testing.T) {
    cache := NewLRUCache(2)

    cache.Set("key1", "value1", 1*time.Second)
    time.Sleep(2 * time.Second)

    if _, ok := cache.Get("key1"); ok {
        t.Error("expected key1 to be expired")
    }
}

func TestLRUCache_Delete(t *testing.T) {
    cache := NewLRUCache(2)

    cache.Set("key1", "value1", 1*time.Minute)
    cache.Delete("key1")

    if _, ok := cache.Get("key1"); ok {
        t.Error("expected key1 to be deleted")
    }
}
