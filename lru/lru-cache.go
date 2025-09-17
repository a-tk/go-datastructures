package lru

import "container/list"

type container[K comparable, V any] struct {
	k K
	v V
}

type Cache[K comparable, V any] struct {
	cache map[K]*list.Element
	list  *list.List
	size  int
}

func New[K comparable, V any](size int) *Cache[K, V] {
	return &Cache[K, V]{
		cache: make(map[K]*list.Element),
		list:  list.New(),
		size:  size,
	}
}

func (c *Cache[K, V]) Get(k K) (v V, found bool) {
	e, ok := c.cache[k]
	if ok {
		c.list.MoveToFront(e)
		return e.Value.(container[K, V]).v, true
	} else {
		return v, false
	}
}

func (c *Cache[K, V]) Put(k K, v V) (stale V, evicted bool) {

	if c.list.Len() == c.size {
		// evict stale
		e := c.list.Back()
		c.list.Remove(e)
		staleCont := e.Value.(container[K, V])
		stale = staleCont.v
		delete(c.cache, staleCont.k)
		evicted = true
	}

	e := c.list.PushFront(container[K, V]{k: k, v: v})
	c.cache[k] = e
	return stale, evicted
}
