package main

import (
	"fmt"
	"time"

	"caching/cache"
)

func inMemoryCache(){
    inMemoryCache := cache.NewInMemoryCache(100)
    defer inMemoryCache.InMemory.StopCleanup()
    inMemoryCache.Set("foo", "bar", 5*time.Minute)
    value, err := inMemoryCache.Get("foo")
    if err == nil {
        fmt.Println("In-Memory Cache:", value)
    } else {
        fmt.Println("Error:", err)
    }
}

func redisCache(){
    redisCache := cache.NewRedisCache("localhost:6379", "", 0)
    redisCache.Set("foo", "bar", 5*time.Minute)
    value, err := redisCache.Get("foo")
    if err == nil {
        fmt.Println("Redis Cache:", value)
    } else {
        fmt.Println("Error:", err)
    }
}

func main() {
    // In-Memory Cache
    inMemoryCache()

    // Redis Cache
    redisCache()
}
