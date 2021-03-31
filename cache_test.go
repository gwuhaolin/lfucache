package lfucache

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testGetSet(t *testing.T, cacheBuilder func(capacity uint) Cache) {
	cache := cacheBuilder(1024)
	_, has := cache.Get("a")
	assert.Equal(t, has, false)

	cache.Set("a", "A")
	v, has := cache.Get("a")
	assert.Equal(t, has, true)
	assert.Equal(t, v, "A")

	cache.Del("a")
	v, has = cache.Get("a")
	assert.Nil(t, v)
	assert.Equal(t, has, false)
}

func testCapacity(t *testing.T, cacheBuilder func(capacity uint) Cache) {
	cache := cacheBuilder(3)
	cache.Set("a", "A")
	cache.Get("a")
	cache.Set("b", "B")
	cache.Get("b")
	assert.Equal(t, cache.Len(), 2)

	cache.Set("c", "C")
	cache.Get("c") // 让c成为访问次数最多的
	cache.Get("c") // 让c成为访问次数最多的
	assert.Equal(t, cache.Len(), 3)

	cache.Set("d", "D")
	assert.Greater(t, 3, cache.Len())

	v, has := cache.Get("c")
	assert.Equal(t, has, true)
	assert.Equal(t, v, "C")
}

func testOOM(t *testing.T, cacheBuilder func(capacity uint) Cache) {
	cache := cacheBuilder(64)
	c := make(chan interface{}, 4096)
	for i := 0; i < 999999; i++ {
		c <- nil
		go func() {
			key := fmt.Sprintf(`%d`, rand.Int())
			cache.Set(key, key)
			t.Log(cache.Len(), key)
			assert.GreaterOrEqual(t, 64, cache.Len())
			<-c
		}()
	}
	<-c
}

func TestFifo(t *testing.T) {
	testGetSet(t, NewFifoCache)
	testCapacity(t, NewFifoCache)
	testOOM(t, NewFifoCache)
}

func TestLFU(t *testing.T) {
	testGetSet(t, NewLfuCache)
	testCapacity(t, NewLfuCache)
	testOOM(t, NewFifoCache)
}
