package lfucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testGetSet(t *testing.T, cacheBuilder func(capacity int) Cache) {
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

func testCapacity(t *testing.T, cacheBuilder func(capacity int) Cache) {
	cache := cacheBuilder(3)
	cache.Set("a", "A")
	cache.Set("b", "B")
	assert.Equal(t, cache.Len(), 2)
	cache.Set("c", "C")
	assert.Equal(t, cache.Len(), 3)
	cache.Set("d", "D")
	assert.Equal(t, cache.Len(), 3)
	v, has := cache.Get("d")
	assert.Equal(t, has, true)
	assert.Equal(t, v, "D")
}

func TestFifo(t *testing.T) {
	testGetSet(t, NewFifoCache)
	testCapacity(t, NewFifoCache)
}

func TestLFU(t *testing.T) {
	testGetSet(t, NewLfuCache)
	testCapacity(t, NewLfuCache)
}
