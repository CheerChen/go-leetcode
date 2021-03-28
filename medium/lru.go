package medium

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
