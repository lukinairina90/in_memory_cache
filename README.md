# In Memory cache library
___

This library contains two implementations of memory cache in two different packages.

Package `cache` contains standard in memory cache implementations with `string` keys and `interface{}` value.

Package `generic_cache` contains in memory cache implementation using golang parameter types (generics).

---

### `cache` package example:

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