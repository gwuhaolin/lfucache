package lfucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testGetSet(t *testing.T, cacheBuilder func(capacity int) Cache) {
	cache := cacheBuilder(1024)
	_, has := cache.Get("a")
	assert.Equal(t, has, false)

	cache.Set("a", 1)
	v, has := cache.Get("a")
	assert.Equal(t, has, true)
	assert.Equal(t, v, 1)
}

func testCapacity(t *testing.T, cacheBuilder func(capacity int) Cache) {
	cache := cacheBuilder(3)
	cache.Set("a", 1)
	cache.Set("b", 2)
	assert.Equal(t, cache.Len(), 2)
	cache.Set("c", 3)
	assert.Equal(t, cache.Len(), 3)
	cache.Set("d", 4)
	assert.Equal(t, cache.Len(), 3)
	v, has := cache.Get("d")
	assert.Equal(t, has, true)
	assert.Equal(t, v, 4)
}

func TestFifo(t *testing.T) {
	testGetSet(t, NewFifoCache)
	testCapacity(t, NewFifoCache)
}

func TestLFU(t *testing.T) {
	testGetSet(t, NewLfuCache)
	testCapacity(t, NewLfuCache)
}
