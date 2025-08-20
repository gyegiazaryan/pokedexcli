package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache() Cache {
	return Cache{
		cache: make(map[string]cacheEntry),
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)
	for key, v := range c.cache {
		if v.createdAt.Before(timeAgo) {
			delete(c.cache, key)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}
