package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	Entries 	map[string]cacheEntry
	CacheMutex	*sync.Mutex
	Ticker		*time.Ticker
	interval	time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		map[string]cacheEntry{},
		&sync.Mutex{},
		time.NewTicker(interval),
		interval,
	}

	go newCache.reapLoop()

	return newCache
}

func (c* Cache) Add(key string, val []byte) {
	fmt.Println("Adding to cache:", key)
	c.CacheMutex.Lock()
	defer c.CacheMutex.Unlock()

	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
}

func (c* Cache) Get(key string) ([]byte, bool) {
	c.CacheMutex.Lock()
	defer c.CacheMutex.Unlock()

	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}

	return entry.value, true
}

func (c* Cache) reapLoop() {
	for range c.Ticker.C {
		for key, entry := range c.Entries {
			if time.Since(entry.createdAt) > c.interval {
				c.CacheMutex.Lock()
				delete(c.Entries, key)
				c.CacheMutex.Unlock()
			}
		}
	}
}