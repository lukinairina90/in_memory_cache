package generic_cache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrKeyNotFound = errors.New("key not found")

type item[V any] struct {
	value   V
	ttl     time.Duration
	setTime time.Time
}

type Cache[K comparable, V any] struct {
	m  map[K]item[V]
	mu sync.RWMutex
}

func New[K comparable, V any]() *Cache[K, V] {
	cache := &Cache[K, V]{m: make(map[K]item[V])}
	cache.runTTLHandler()
	return cache
}

func (c *Cache[K, V]) Set(key K, value V, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = item[V]{
		value:   value,
		ttl:     ttl,
		setTime: time.Now(),
	}
	return nil
}

func (c *Cache[K, V]) Get(key K) (V, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.m[key]
	if !ok {
		return val.value, ErrKeyNotFound
	}
	return val.value, nil
}

func (c *Cache[K, V]) Delete(key K) error {
	return c.delete(key)
}

func (c *Cache[K, V]) delete(key K) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.m[key]; !ok {
		return ErrKeyNotFound
	}

	delete(c.m, key)

	return nil
}

func (c *Cache[K, V]) runTTLHandler() {
	go func() {
		t := time.NewTicker(time.Millisecond * 100)
		for range t.C {
			for key, val := range c.m {
				if val.setTime.Add(val.ttl).Before(time.Now()) {
					if err := c.delete(key); err != nil {
						fmt.Printf("error deleting key [%s] with error: %s\n", key, err)
					}
				}
			}
		}
	}()
}
