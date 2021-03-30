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
	cache.Set("b", "B")
	assert.Equal(t, cache.Len(), 2)
	cache.Set("c", "C")
	assert.Equal(t, cache.Len(), 3)
	cache.Set("d", "D")
	assert.Greater(t, 3, cache.Len())
	v, has := cache.Get("d")
	assert.Equal(t, has, true)
	assert.Equal(t, v, "D")
}

func TestFifo(t *testing.T) {
	testGetSet(t, NewFifoCache)
	testCapacity(t, NewFifoCache)
	//testOOM(t, NewFifoCache)
}

func testOOM(t *testing.T, cacheBuilder func(capacity uint) Cache) {
	cache := cacheBuilder(10240)
	c := make(chan interface{}, 1000)
	for {
		c <- nil
		go func() {
			key := fmt.Sprintf(`%d`, rand.Int())
			cache.Set(key, key)
			t.Log(cache.Len(), key)
			<-c
		}()
	}
	<-c
}

//func TestLFU(t *testing.T) {
//	testGetSet(t, NewLfuCache)
//	testCapacity(t, NewLfuCache)
//}
