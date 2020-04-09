A simple in memory LFU(Least Frequently Used) cache with high-performance and concurrency support for golang

```go
cache := lfucache.NewLfuCache(1024) // max cache 1024 items 

// read
if val, has := cache.Get("key"); has {
    // use val
    log.Print(val)
}

// write
cache.Set("key", val)

// clear all cache
cache.Clear()
```