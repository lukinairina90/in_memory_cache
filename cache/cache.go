package cache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrKeyNotFound = errors.New("key not found")

type item struct {
	value   interface{}
	ttl     time.Duration
	setTime time.Time
}

type Cache struct {
	m  map[string]item
	mu sync.RWMutex
}

func New() *Cache {
	cache := &Cache{m: make(map[string]item)}
	cache.runTTLHandler()

	return cache
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = item{
		value:   value,
		ttl:     ttl,
		setTime: time.Now(),
	}
	return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.m[key]
	if !ok {
		return nil, ErrKeyNotFound
	}
	return val.value, nil
}

func (c *Cache) Delete(key string) error {
	return c.delete(key)
}

func (c *Cache) delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.m[key]; !ok {
		return ErrKeyNotFound
	}

	delete(c.m, key)

	return nil
}

func (c *Cache) runTTLHandler() {
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
