package main

import (
	"container/list"
	"sync"
)

type LRU[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Set(key K, value V)
}
type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	mu       sync.Mutex
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	var cache LRUCache[K, V]
	cache.capacity = capacity
	cache.ll.Init()
	cache.items = make(map[K]*list.Element)
	return &cache
}
func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	elem, ok := c.items[key]
	if !ok {
		return
	}
	value = elem.Value.(entry[K, V]).value
	c.ll.MoveToBack(elem)
	return value, ok
}
func (c *LRUCache[K, V]) Set(key K, value V) {
	// Если емкость 0 - ничего не делаем
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if c.capacity <= 0 {
		return
	}

	// проверяем наличие значения в мапе и если такой ключ уже есть - обновляем значение
	elem, ok := c.items[key]
	if ok {

		elem.Value = entry[K,V]{key: key, value: value}
		return
	}

	// если кеш еще не заполнен и значение новое
	if c.ll.Len() < c.capacity {
		elem := c.ll.PushBack(entry[K,V]{key: key, value: value})
		c.items[key] = elem
		return
	}

	// если кеш заполнен и значение новое
	// получаем самое старое значение и удаляем его из списка и мапы
	oldestValueElem := c.ll.Front()
	delete(c.items, oldestValueElem.Value.(entry[K,V]).key)
	c.ll.Remove(oldestValueElem)

	// Записываем новые значения
	elem = c.ll.PushBack(entry[K,V]{key: key, value: value})
	c.items[key] = elem
}
