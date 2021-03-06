package lfucache

import (
	"sync"
)

type FifoCache struct {
	Cache
	Capacity uint
	valueMap map[string]*index
	lock     *sync.RWMutex
	nowIndex uint
}

type index struct {
	value interface{}
	flag  uint
}

//const maxUint = ^uint(0)

func NewFifoCache(capacity uint) Cache {
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
	c.lock.Lock()
	defer c.lock.Unlock()
	f, has := c.valueMap[key]
	if has {
		f.value = value
	} else {
		// 出现 constant 18446744073709551616 overflows uint 的概览太小了，先忽略这个逻辑
		//if c.nowIndex >= maxUint {
		//	// 清除
		//	c.valueMap = map[string]*index{}
		//	c.nowIndex = 0
		//}
		c.nowIndex++
		c.valueMap[key] = &index{
			value: value,
			flag:  c.nowIndex,
		}
		// 清理最老的1/2
		if uint(len(c.valueMap)) > c.Capacity {
			min := c.nowIndex - c.Capacity/2
			for k, f := range c.valueMap {
				if f.flag < min {
					delete(c.valueMap, k)
				}
			}
		}
	}
}

func (c *FifoCache) Del(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	_, has := c.valueMap[key]
	if !has {
		return
	}
	delete(c.valueMap, key)
}

// 清空
func (c *FifoCache) Clear() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	count := len(c.valueMap)
	c.valueMap = map[string]*index{}
	c.nowIndex = 0
	return count
}

// length of current cache list
func (c *FifoCache) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.valueMap)
}
