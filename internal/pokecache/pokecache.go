package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt   time.Time
	val         []byte
	nextUrl     string
	previousUrl string
}

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
	}

	go cache.expiryLoop(interval)

	return &cache
}

func (c *Cache) expiryLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()

		for k, v := range c.entries {
			elapsed := time.Since(v.createdAt)
			if elapsed > interval {
				delete(c.entries, k)
			}
		}

		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte, nextUrl string, previousUrl string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt:   time.Now(),
		val:         val,
		nextUrl:     nextUrl,
		previousUrl: previousUrl,
	}
}

func (c *Cache) Get(key string) ([]byte, string, string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, "", "", false
	}
	return entry.val, entry.nextUrl, entry.previousUrl, true
}
