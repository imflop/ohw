package hw04_lru_cache //nolint:golint,stylecheck

import "github.com/imflop/ohw/hw04_lru_cache/list"

type Key string

type Cache interface {
	// Place your code here
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Clear()
}

type lruCache struct {
	// Place your code here:
	// - capacity
	// - queue
	// - items
	capacity int
	queue *list.List
	items map[int]*list.listItem
}

type cacheItem struct {
	// Place your code here
	key string
	value int
}

func NewCache(capacity int) Cache {
	return &lruCache{}
}
