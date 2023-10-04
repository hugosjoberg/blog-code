package cache

import "sync"

type Cache struct {
	mu    sync.RWMutex
	cache map[string]string
}

func (c *Cache) Set(key string, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

func (c *Cache) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, _ := c.cache[key]
	return value
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]string),
	}
}
