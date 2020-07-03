package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    *list
	items    map[string]*listItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    new(list),
		items:    make(map[string]*listItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if n, ok := l.items[string(key)]; ok {
		l.queue.Remove(n)
		n.Value = cacheItem{key: key, value: value}
		l.queue.PushFront(n.Value)
		return true
	}
	if l.queue.Len() == l.capacity {
		k := l.queue.Back().Value.(*listItem).Value.(cacheItem).key
		l.queue.Remove(l.queue.Back())
		delete(l.items, string(k))
	}
	n := cacheItem{
		key:   key,
		value: value,
	}
	p := l.queue.PushFront(n)
	l.items[string(key)] = p
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if n, ok := l.items[string(key)]; ok {
		v := n.Value.(cacheItem).value
		l.queue.MoveToFront(n)
		return v, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	for l.capacity != 0 {
		v := l.queue.Back()
		l.queue.Remove(v)
		delete(l.items, string(v.Value.(cacheItem).key))
		l.capacity--
	}
}
