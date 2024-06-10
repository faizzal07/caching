package cache

import (
	"sync"
	"time"
)

type LRUCache struct {
    maxSize   int
    ll        *LinkedList
    cache     map[string]*ListNode
    mutex     sync.Mutex
    stopClean chan struct{}
}

func NewLRUCache(maxSize int) *LRUCache {
    c := &LRUCache{
        maxSize:   maxSize,
        ll:        NewLinkedList(),
        cache:     make(map[string]*ListNode),
        stopClean: make(chan struct{}),
    }
    go c.cleanupExpiredEntries()
    return c
}

func (c *LRUCache) Set(key string, value interface{}, ttl time.Duration) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if node, ok := c.cache[key]; ok {
        c.ll.MoveToFront(node)
        node.value = value
        node.ttl = time.Now().Add(ttl)
        return
    }

    node := &ListNode{key: key, value: value, ttl: time.Now().Add(ttl)}
    c.ll.PushFront(node)
    c.cache[key] = node

    if len(c.cache) > c.maxSize {
        c.removeOldest()
    }
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if node, ok := c.cache[key]; ok {
        if time.Now().After(node.ttl) {
            c.removeNode(node)
            return nil, false
        }
        c.ll.MoveToFront(node)
        return node.value, true
    }
    return nil, false
}

func (c *LRUCache) Delete(key string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if node, ok := c.cache[key]; ok {
        c.removeNode(node)
    }
}

func (c *LRUCache) removeOldest() {
    node := c.ll.RemoveLast()
    if node != nil {
        c.removeNode(node)
    }
}

func (c *LRUCache) removeNode(node *ListNode) {
    c.ll.remove(node)
    delete(c.cache, node.key)
}

func (c *LRUCache) cleanupExpiredEntries() {
    ticker := time.NewTicker(time.Minute)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            c.mutex.Lock()
            for _, node := range c.cache {
                if time.Now().After(node.ttl) {
                    c.removeNode(node)
                }
            }
            c.mutex.Unlock()
        case <-c.stopClean:
            return
        }
    }
}

func (c *LRUCache) StopCleanup() {
    close(c.stopClean)
}
