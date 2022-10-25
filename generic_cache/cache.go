package generic_cache

type Cache[K comparable, V any] struct {
	m map[K]V
}

func New[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{m: make(map[K]V)}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.m[key] = value
}

func (c *Cache[K, V]) Get(key K) V {
	return c.m[key]
}

func (c *Cache[K, V]) Delete(key K) {
	delete(c.m, key)
}
