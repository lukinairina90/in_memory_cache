# In Memory cache library
___

This library contains two implementations of memory cache in two different packages.

Package `cache` contains standard in memory cache implementations with `string` keys and `interface{}` value.

Package `generic_cache` contains in memory cache implementation using golang parameter types (generics).

---

### `cache` package example:

```go
import (
	"fmt"
	
	"github.com/lukinairina90/in_memory_cache/cache"
)

func main() {
	mc := cache.New()
	mc.Set("user-id", 42)
	uID := mc.Get("user-id")
	fmt.Println(uID)
	
	mc.Delete("user-id")
	uID = mc.Get("user-id")
	fmt.Println(uID)
}
```

---
### `generic_cache` package example:

```go
import (
	"fmt"
	
	"github.com/lukinairina90/in_memory_cache/generic_cache"
)

func main() {
	mc := generic_cache.New[string, int]() // pass parameter types for your generic cache object. 
	mc.Set("user-id", 42)
	uID := mc.Get("user-id")
	fmt.Println(uID)

	mc.Delete("user-id")
	uID = mc.Get("user-id")
	fmt.Println(uID)
}
```