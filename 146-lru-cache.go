package main

//运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
//获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

import "container/list"

type LRUCache struct {
	hhash map[int]*list.Element
	llist *list.List
	size  int
}

type entry struct {
	key   int
	value int
}

// func Constructor(capacity int) LRUCache {
// 	var lc LRUCache
// 	lc.hhash = make(map[int]*list.Element)
// 	lc.llist = list.New()
// 	lc.size = capacity
// 	return lc
// }

func (cache *LRUCache) Get(key int) int {
	if e, ok := cache.hhash[key]; ok {
		cache.llist.MoveToFront(e)
		return e.Value.(entry).value
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if e, ok := cache.hhash[key]; ok {
		e.Value = entry{key, value}
		cache.llist.MoveToFront(e)
	} else {
		node := entry{key, value}
		cache.hhash[node.key] = cache.llist.PushFront(node)
		if cache.llist.Len() > cache.size {
			tail := cache.llist.Back()
			if tail != nil {
				cache.llist.Remove(tail)
				delete(cache.hhash, tail.Value.(entry).key)
			}
		}
	}
}
