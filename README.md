![Go](https://github.com/gwuhaolin/lfucache/workflows/Go/badge.svg)
A simple in memory cache with high-performance and concurrency support for golang, includes LFU(Least Frequently Used) & FIFO(First In First Out).

```go
import (
 "github.com/gwuhaolin/lfucache"
)

cache := lfucache.NewLfuCache(1024) // max cache 1024 items 
// cache := lfucache.NewFifoCache(1024) // max cache 1024 items 

// read
if val, has := cache.Get("key"); has {
    // use val
    log.Print(val)
}

// write
cache.Set("key", val)

// clear all cache
len:=cache.Clear()

// cache current length
len:=cache.Len()
```