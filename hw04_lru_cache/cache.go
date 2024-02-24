package hw04lrucache

import (
	"fmt"
	"sync"
)

type Key string

type cacheItem struct {
	key   Key
	value interface{}
}

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
	PrintCacheState(stage string)
}

type lruCache struct {
	sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.Lock()
	defer l.Unlock()

	cachedItem := cacheItem{key: key, value: value}

	item, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(item)
		item.Value = cachedItem
		return true
	}

	if l.queue.Len() == l.capacity {
		removable := l.queue.Back()
		l.queue.Remove(removable)
		removableItem := removable.Value.(cacheItem)
		delete(l.items, removableItem.key)
	}

	item = l.queue.PushFront(cachedItem)
	l.items[key] = item
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.Lock()
	defer l.Unlock()

	item, ok := l.items[key]
	if !ok {
		return nil, false
	}

	l.queue.MoveToFront(item)
	origValue := item.Value.(cacheItem).value
	return origValue, true
}

func (l *lruCache) Clear() {
	l.Lock()
	defer l.Unlock()

	for l.queue.Len() > 0 {
		lastItem := l.queue.Back()
		l.queue.Remove(lastItem)
	}

	l.items = make(map[Key]*ListItem)
}

func (l *lruCache) PrintCacheState(stage string) {
	fmt.Printf("%s\n", stage)
	fmt.Println("There is Queue values:")
	prev := l.queue.Back()
	for i := 0; i < 5; i++ {
		if prev == nil {
			break
		}
		fmt.Printf("%v\n", prev)
		prev = prev.Prev
	}

	fmt.Println("\nThere is Cache values:")
	for key, value := range l.items {
		fmt.Printf("Key: %s, Val: %v\n", key, value.Value)
	}
	fmt.Printf("\n")
}
