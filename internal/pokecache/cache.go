package pokecache

import (
	"errors"
	"fmt"
	"io"
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

func (c *Cache) Get(url string) ([]byte, error) {
	defer logResponseTime(time.Now())
	c.mu.Lock()
	defer c.mu.Unlock()
	var rawResponse []byte
	entry, ok := c.entries[url]
	if ok {
		rawResponse = entry.val
	} else {
		res, err := rawFromWeb(url)
		if err != nil {
			return res, err
		}
		rawResponse = res
	}
	c.entries[url] = cacheEntry{
		createdAt: time.Now(),
		val:       rawResponse,
	}

	res, err := responseFromRaw(rawResponse)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return []byte{}, errors.New(res.Status)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}

func (c *Cache) GetRaw(url string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[url]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) Dump() string {
	var b strings.Builder
	for k, entry := range c.entries {
		fmt.Fprintf(&b, "[%s: [createdAt: %v, val: ...]\n", k, entry.createdAt)
	}
	return b.String()
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
