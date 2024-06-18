package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type CacheElement struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	item, ok := lc.items[key]
	element := CacheElement{key, value}

	if ok {
		item.Value = element
		lc.queue.MoveToFront(item)
		return true
	}

	lc.queue.PushFront(element)
	lc.items[key] = lc.queue.Front()
	if lc.queue.Len() > lc.capacity {
		back := lc.queue.Back().Value.(CacheElement)
		lc.queue.Remove(lc.queue.Back())
		delete(lc.items, back.key)
	}
	return false
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := lc.items[key]
	if ok {
		lc.queue.MoveToFront(item)
		return item.Value.(CacheElement).value, true
	}
	return nil, false
}

func (lc *lruCache) Clear() {
	lc = &lruCache{
		capacity: lc.capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, lc.capacity),
	}
}
