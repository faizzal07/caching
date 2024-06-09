package main

import (
	"fmt"
	"time"

	"caching/cache"
)

func main() {
    // In-Memory Cache
    fmt.Println("dfs")
    inMemoryCache := cache.NewInMemoryCache(100)
    defer inMemoryCache.InMemory.StopCleanup()
    inMemoryCache.Set("foo", "bar", 5*time.Minute)
    value, err := inMemoryCache.Get("foo")
    if err == nil {
        fmt.Println("In-Memory Cache:", value)
    } else {
        fmt.Println("Error:", err)
    }

    // Redis Cache
    redisCache := cache.NewRedisCache("localhost:6379", "", 0)
    redisCache.Set("foo", "bar", 5*time.Minute)
    value, err = redisCache.Get("foo")
    if err == nil {
        fmt.Println("Redis Cache:", value)
    } else {
        fmt.Println("Error:", err)
    }

    // Memcached Cache
    memcachedCache := cache.NewMemcachedCache("localhost:11211")
    memcachedCache.Set("foo", []byte("bar"), 5*time.Minute)
    value, err = memcachedCache.Get("foo")
    if err == nil {
        fmt.Println("Memcached Cache:", string(value.([]byte)))
    } else {
        fmt.Println("Error:", err)
    }
}
