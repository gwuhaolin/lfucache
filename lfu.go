package lfucache

import "sync"

type LfuCache struct {
	Cache
	Capacity int
	valueMap map[string]*freq
	lock     *sync.RWMutex
}

type freq struct {
	value interface{}
	flag  int
}

func NewLfuCache(capacity int) Cache {
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
		f.flag++
		return f.value, has
	}
}

func (c *LfuCache) Set(key string, value interface{}) {
	c.lock.RLock()
	f, has := c.valueMap[key]
	c.lock.RUnlock()
	if has {
		f.value = value
	} else {
		c.lock.Lock()
		c.valueMap[key] = &freq{
			value: value,
			flag:  1,
		}
		// 清理访问次数最少的
		if len(c.valueMap) > c.Capacity {
			minKey := key
			minFreq := c.valueMap[minKey]
			for k, f := range c.valueMap {
				if f.flag <= minFreq.flag {
					minKey = k
					minFreq = f
					if f.flag == 1 {
						break
					}
				}
			}
			if minKey != key {
				delete(c.valueMap, minKey)
			}
		}
		c.lock.Unlock()
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

// length of current cache list
func (c *LfuCache) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.valueMap)
}
