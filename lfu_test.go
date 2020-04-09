package lfucache

import "testing"

func TestLfuCache(t *testing.T) {
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
