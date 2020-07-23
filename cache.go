package lfucache

type Cache interface {
	Get(key string) (val interface{}, has bool)
	Set(key string, value interface{})
	Clear() int
	Len() int
}
