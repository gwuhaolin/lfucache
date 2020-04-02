package lfucache

import "sync"

type LfuCache struct {
	Capacity int
	valueMap map[string]*freq
	lock     *sync.RWMutex
}

type freq struct {
	value interface{}
	count int
}

func NewLfuCache(capacity int) *LfuCache {
	return &LfuCache{
		Capacity: capacity,
		valueMap: make(map[string]*freq),
		lock:     new(sync.RWMutex),
	}
}

func (c *LfuCache) Get(key string) (val interface{}, has bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	f, has := c.valueMap[key]
	if !has {
		return nil, has
	} else {
		f.count++
		return f.value, has
	}
}

func (c *LfuCache) Set(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	f, has := c.valueMap[key]
	if has {
		f.value = value
	} else {
		c.valueMap[key] = &freq{
			value: value,
			count: 1,
		}
	}

	// 清理访问次数最少的
	if len(c.valueMap) >= c.Capacity {
		minKey := key
		minFreq := c.valueMap[minKey]
		for k, f := range c.valueMap {
			if f.count <= minFreq.count {
				minKey = k
				minFreq = f
			}
		}
		if minKey != key {
			delete(c.valueMap, minKey)
		}
	}
}

// 清空
func (c *LfuCache) Clear() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	count := len(c.valueMap)
	for k := range c.valueMap {
		delete(c.valueMap, k)
	}
	return count
}
