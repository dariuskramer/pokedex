package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mutex   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
		mutex:   &sync.Mutex{},
	}

	go cache.expiryLoop(interval)

	return cache
}

func (c *Cache) expiryLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()

		for k, v := range c.entries {
			elapsed := time.Since(v.createdAt)
			if elapsed > interval {
				delete(c.entries, k)
			}
		}

		c.mutex.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, ok := c.entries[key]
	return entry.val, ok
}
