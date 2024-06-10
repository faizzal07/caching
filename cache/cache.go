package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type Backend int

const (
    InMemory Backend = iota
    Redis
    Memcached
)

type Cache struct {
    InMemory  *LRUCache
    redis     *redis.Client
    backend   Backend
    ctx       context.Context
}

func NewInMemoryCache(size int) *Cache {
    return &Cache{
        InMemory: NewLRUCache(size),
        backend:  InMemory,
    }
}

func NewRedisCache(addr string, password string, db int) *Cache {
    return &Cache{
        redis: redis.NewClient(&redis.Options{
            Addr:     addr,
            Password: password,
            DB:       db,
        }),
        backend: Redis,
        ctx:     context.Background(),
    }
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error {
    switch c.backend {
    case InMemory:
        c.InMemory.Set(key, value, ttl)
    case Redis:
        return c.redis.Set(c.ctx, key, value, ttl).Err()
    default:
        return errors.New("unsupported backend")
    }
    return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
    switch c.backend {
    case InMemory:
        value, ok := c.InMemory.Get(key)
        if !ok {
            return nil, errors.New("key not found")
        }
        return value, nil
    case Redis:
        return c.redis.Get(c.ctx, key).Result()
    default:
        return nil, errors.New("unsupported backend")
    }
}

func (c *Cache) Delete(key string) error {
    switch c.backend {
    case InMemory:
        c.InMemory.Delete(key)
    case Redis:
        return c.redis.Del(c.ctx, key).Err()
    default:
        return errors.New("unsupported backend")
    }
    return nil
}
