package cache

import (
	"testing"
	"time"
)

func TestInMemoryCache_SetAndGet(t *testing.T) {
    cache := NewInMemoryCache(2)

    cache.Set("key1", "value1", 1*time.Minute)
    if value, err := cache.Get("key1"); err != nil || value != "value1" {
        t.Errorf("expected value1, got %v", value)
    }
}

func TestRedisCache_SetAndGet(t *testing.T) {
    cache := NewRedisCache("localhost:6379", "", 0)

    cache.Set("key1", "value1", 1*time.Minute)
    if value, err := cache.Get("key1"); err != nil || value != "value1" {
        t.Errorf("expected value1, got %v", value)
    }
}

func TestMemcachedCache_SetAndGet(t *testing.T) {
    cache := NewMemcachedCache("localhost:11211")

    cache.Set("key1", []byte("value1"), 1*time.Minute)
    if value, err := cache.Get("key1"); err != nil || string(value.([]byte)) != "value1" {
        t.Errorf("expected value1, got %v", value)
    }
}
