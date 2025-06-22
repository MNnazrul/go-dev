package cache

import (
	"sync"
	"time"
)

type CachedResponse struct {
	StatusCode int
	Header     map[string][]string
	Body       []byte
	CreatedAt  time.Time
}

type MemoryCache struct {
	mu    sync.RWMutex
	store map[string]CachedResponse
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		store: make(map[string]CachedResponse),
	}
}

func (c *MemoryCache) Get(key string) (CachedResponse, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	resp, ok := c.store[key]
	return resp, ok
}

func (c *MemoryCache) Set(key string, resp CachedResponse) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = resp
}

func (c *MemoryCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store = make(map[string]CachedResponse)
}
