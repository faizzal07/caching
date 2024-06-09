package cache

import (
	"container/list"
	"sync"
	"time"
)

type LRUCache struct {
    maxSize   int
    ll        *list.List
    cache     map[string]*list.Element
    mutex     sync.Mutex
    stopClean chan struct{}
}

type entry struct {
    key   string
    value interface{}
    ttl   time.Time
}

func NewLRUCache(maxSize int) *LRUCache {
    c := &LRUCache{
        maxSize:   maxSize,
        ll:        list.New(),
        cache:     make(map[string]*list.Element),
        stopClean: make(chan struct{}),
    }
    go c.cleanupExpiredEntries()
    return c
}

func (c *LRUCache) Set(key string, value interface{}, ttl time.Duration) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if elem, ok := c.cache[key]; ok {
        c.ll.MoveToFront(elem)
        elem.Value.(*entry).value = value
        elem.Value.(*entry).ttl = time.Now().Add(ttl)
        return
    }

    elem := c.ll.PushFront(&entry{key, value, time.Now().Add(ttl)})
    c.cache[key] = elem

    if c.ll.Len() > c.maxSize {
        c.removeOldest()
    }
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if elem, ok := c.cache[key]; ok {
        if time.Now().After(elem.Value.(*entry).ttl) {
            c.removeElement(elem)
            return nil, false
        }
        c.ll.MoveToFront(elem)
        return elem.Value.(*entry).value, true
    }
    return nil, false
}

func (c *LRUCache) Delete(key string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if elem, ok := c.cache[key]; ok {
        c.removeElement(elem)
    }
}

func (c *LRUCache) removeOldest() {
    elem := c.ll.Back()
    if elem != nil {
        c.removeElement(elem)
    }
}

func (c *LRUCache) removeElement(elem *list.Element) {
    c.ll.Remove(elem)
    delete(c.cache, elem.Value.(*entry).key)
}

func (c *LRUCache) cleanupExpiredEntries() {
    ticker := time.NewTicker(time.Minute)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            c.mutex.Lock()
            for _, elem := range c.cache {
                if time.Now().After(elem.Value.(*entry).ttl) {
                    c.removeElement(elem)
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
