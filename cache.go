package lfucache

type Cache interface {
	Get(key string) (val interface{}, has bool)
	OptGet(key string) (val interface{}, has bool)
	Set(key string, value interface{})
	Del(key string)
	Clear() int
	Len() int
}
