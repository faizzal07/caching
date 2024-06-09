package cache

import (
	"testing"
	"time"
)

func BenchmarkLRUCache_Set(b *testing.B) {
    cache := NewLRUCache(1000)
    for i := 0; i < b.N; i++ {
        cache.Set("key"+string(i), "value", 1*time.Minute)
    }
}

func BenchmarkLRUCache_Get(b *testing.B) {
    cache := NewLRUCache(1000)
    cache.Set("key", "value", 1*time.Minute)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        cache.Get("key")
    }
}

func BenchmarkLRUCache_Delete(b *testing.B) {
    cache := NewLRUCache(1000)
    cache.Set("key", "value", 1*time.Minute)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        cache.Delete("key")
    }
}
