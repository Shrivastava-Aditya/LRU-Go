package main

import (
	"container/list"
	"sync"
	"time"
)

// CacheItem represents an item in the cache
type CacheItem struct {
	Key        string
	Value      interface{}
	Expiration time.Time
}

// LRUCache represents the LRU cache
type LRUCache struct {
	cache   map[string]*list.Element
	lruList *list.List
	maxKeys int
	mutex   sync.Mutex
}

// NewLRUCache creates a new LRUCache instance
func NewLRUCache(maxKeys int) *LRUCache {
	return &LRUCache{
		cache:   make(map[string]*list.Element),
		lruList: list.New(),
		maxKeys: maxKeys,
	}
}

// Get retrieves a value from the cache given a key
func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		item := elem.Value.(*CacheItem)
		if item.Expiration.After(time.Now()) {
			c.lruList.MoveToFront(elem)
			return item.Value, true
		}
		// If the item has expired, remove it from cache
		c.removeElement(elem)
	}
	return nil, false
}

// Set sets a key-value pair in the cache with expiration
func (c *LRUCache) Set(key string, value interface{}, expiration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		c.lruList.MoveToFront(elem)
		item := elem.Value.(*CacheItem)
		item.Value = value
		item.Expiration = time.Now().Add(expiration)
	} else {
		if len(c.cache) >= c.maxKeys {
			c.evictOldest()
		}
		item := &CacheItem{
			Key:        key,
			Value:      value,
			Expiration: time.Now().Add(expiration),
		}
		elem := c.lruList.PushFront(item)
		c.cache[key] = elem
	}
}

// removeElement removes an element from the cache
func (c *LRUCache) removeElement(elem *list.Element) {
	c.lruList.Remove(elem)
	delete(c.cache, elem.Value.(*CacheItem).Key)
}

// evictOldest evicts the oldest item from the cache
func (c *LRUCache) evictOldest() {
	elem := c.lruList.Back()
	if elem != nil {
		c.removeElement(elem)
	}
}
