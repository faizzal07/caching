package cache

import (
	"testing"
	"time"
)

func TestInMemoryCache_SetAndGet(t *testing.T) {
    cache := NewInMemoryCache(2)

    err := cache.Set("key1", "value1", 1*time.Minute)
    if err != nil {
        t.Errorf("failed to set value: %v", err)
    }
    value, err := cache.Get("key1")
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    if value != "value1" {
        t.Errorf("expected value1, got %v", value)
    }
}

func TestRedisCache_SetAndGet(t *testing.T) {
    cache := NewRedisCache("localhost:6379", "", 0)

    err := cache.Set("key1", "value1", 1*time.Minute)
    if err != nil {
        t.Errorf("failed to set value: %v", err)
    }

    value, err := cache.Get("key1")
    if err != nil {
        t.Errorf("failed to get value: %v", err)
    }
	if value != "value1" {
        t.Errorf("expected value1, got %v", value)
    }
}
