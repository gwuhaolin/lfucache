package lfucache

import "testing"

func TestGetSet(t *testing.T) {
	cache := NewLfuCache(1024)
	v, has := cache.Get("a")
	if v != nil || has {
		t.Error(v, has)
	}

	cache.Set("a", 1)
	v, has = cache.Get("a")
	if v == nil || !has {
		t.Error(v, has)
	}
}

func TestCapacity(t *testing.T) {
	cache := NewLfuCache(3)
	cache.Set("a", 1)
	cache.Set("b", 1)
	if cache.Len() != 2 {
		t.Error(cache.Len())
	}
	cache.Set("c", 1)
	if cache.Len() != 3 {
		t.Error(cache.Len())
	}
	cache.Set("d", 1)
	if cache.Len() != 3 {
		t.Error(cache.Len())
	}
	v, has := cache.Get("a")
	if v != nil || has {
		t.Error(v, has)
	}
}
