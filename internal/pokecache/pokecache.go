package pokecache

import (
	"time"
)

type Cache struct {
	//mutex sync.Mutex
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	//performing the purging in a seperate thread or else the code gets stuck here.
	go c.purgeLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	//c.mutex.Lock()
	//defer c.mutex.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	return entry.val, ok
}

// deletes any cache records that are older than delete interval
func (c *Cache) PurgeCache(deleteInterval time.Duration) {
	deleteThreshold := time.Now().UTC().Add(-deleteInterval)
	for key, v := range c.cache {
		if v.createdAt.Before(deleteThreshold) {
			delete(c.cache, key)
		}
	}
}

func (c *Cache) purgeLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.PurgeCache(interval)
	}
}
