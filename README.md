# In Memory cache library
___

This library contains two implementations of memory cache in two different packages.

Package `cache` contains standard in memory cache implementations with `string` keys and `interface{}` value.

Package `generic_cache` contains in memory cache implementation using golang parameter types (generics).

---

### `cache` package example:

#### When calling the Set(key string, value interface{}, ttl time.Duration) method,
#### an additional argument is passed for a ttl of the time.Duration type, after which the value will be cleared from the cache.

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/lukinairina90/in_memory_cache/cache"
)

func main() {
	mc := cache.New()
	if err := mc.Set("userId", 42, time.Second*5); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 3)

	m, err := mc.Get("userId")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}

```

---
### `generic_cache` package example:

#### When calling the Set(key comparable, value any, ttl time.Duration) method,
#### an additional argument is passed for a ttl of the time.Duration type, after which the value will be cleared from the cache.

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/lukinairina90/in_memory_cache/generic_cache"

)

func main() {
	mc := generic_cache.New[string, int]()
	if err := mc.Set("userId", 42, time.Second*5); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 3)

	m, err := mc.Get("userId")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}

```