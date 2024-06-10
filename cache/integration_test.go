package cache

import (
	"testing"
	"time"

	"context"

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
