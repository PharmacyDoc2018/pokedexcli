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
	cacheMap map[string]cacheEntry
	mu       sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func NewCache(interval time.Duration) *Cache {
	var cache Cache
	go cacheRun(&cache, interval)
	return &cache
}

func cacheRun(cache *Cache, interval time.Duration) {
	//
}
