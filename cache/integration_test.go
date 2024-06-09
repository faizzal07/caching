package cache

import (
	"testing"
	"time"

	"context"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis/v8"
)

func TestRedisIntegration(t *testing.T) {
    ctx := context.Background()
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    defer rdb.Close()

    err := rdb.Set(ctx, "key1", "value1", 1*time.Minute).Err()
    if err != nil {
        t.Fatalf("Failed to set value in Redis: %v", err)
    }

    val, err := rdb.Get(ctx, "key1").Result()
    if err != nil {
        t.Fatalf("Failed to get value from Redis: %v", err)
    }
    if val != "value1" {
        t.Errorf("Expected value1, got %v", val)
    }
}

func TestMemcachedIntegration(t *testing.T) {
    mc := memcache.New("localhost:11211")
    defer mc.DeleteAll()

    err := mc.Set(&memcache.Item{Key: "key1", Value: []byte("value1"), Expiration: 60})
    if err != nil {
        t.Fatalf("Failed to set value in Memcached: %v", err)
    }

    item, err := mc.Get("key1")
    if err != nil {
        t.Fatalf("Failed to get value from Memcached: %v", err)
    }
    if string(item.Value) != "value1" {
        t.Errorf("Expected value1, got %v", string(item.Value))
    }
}
