package pokecache

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

func NewCache(duration ...time.Duration) *Cache {
	d := 5 * time.Minute
	if len(duration) > 0 {
		d = duration[0]
	}

	nc := &Cache{entries: map[string]cacheEntry{}, mu: &sync.Mutex{}}
	ch := make(chan any)
	go nc.reapLoop(d, ch)
	go nc.prune(d, ch)

	return nc
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	if ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(duration time.Duration, ch chan any) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ch <- c
		}
	}
}

func (c *Cache) prune(duration time.Duration, ch chan any) {
	for {
		select {
		case <-ch:
			t := time.Now()
			c.mu.Lock()
			for i, entry := range c.entries {
				if t.Sub(entry.createdAt) > duration {
					delete(c.entries, i)
				}
			}
			c.mu.Unlock()
		}
	}
}

func (c *Cache) Dump() string {
	var b strings.Builder
	for k, entry := range c.entries {
		fmt.Fprintf(&b, "[%s: [createdAt: %v, val: ...]", k, entry.createdAt)
	}
	return b.String()
}
