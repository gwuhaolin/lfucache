package lfucache

import "sync"

type FifoCache struct {
	Cache
	Capacity int
	valueMap map[string]*index
	lock     *sync.RWMutex
	nowIndex uint
}

type index struct {
	value interface{}
	i     uint
}

const maxUint = ^uint(0)

func NewFifoCache(capacity int) *FifoCache {
	return &FifoCache{
		Capacity: capacity,
		valueMap: make(map[string]*index),
		lock:     new(sync.RWMutex),
	}
}

func (c *FifoCache) Get(key string) (val interface{}, has bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	f, has := c.valueMap[key]
	if !has {
		return nil, has
	} else {
		return f.value, has
	}
}

func (c *FifoCache) Set(key string, value interface{}) {
	c.lock.RLock()
	f, has := c.valueMap[key]
	c.lock.RUnlock()
	if has {
		f.value = value
	} else {
		if c.nowIndex > maxUint {
			c.Clear()
			return
		}
		c.lock.Lock()
		c.nowIndex++
		c.valueMap[key] = &index{
			value: value,
			i:     c.nowIndex,
		}
		// 清理访问次数最少的
		if len(c.valueMap) > c.Capacity {
			minKey := key
			minI := c.valueMap[minKey].i
			for k, f := range c.valueMap {
				if f.i <= minI {
					minKey = k
					minI = f.i
				}
			}
			delete(c.valueMap, minKey)
		}
		c.lock.Unlock()
	}
}

// 清空
func (c *FifoCache) Clear() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	count := len(c.valueMap)
	for k := range c.valueMap {
		delete(c.valueMap, k)
	}
	return count
}

// length of current cache list
func (c *FifoCache) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.valueMap)
}
