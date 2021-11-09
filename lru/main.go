package main

import (
	"container/list"
	"sync"
	"sync/atomic"
)

//思想：最近使用原则，如果满了剔除最久的一个元素。主要用双向链表和map实现
//最后需要一个状态统计命中数和未命中数
type CacheStatus struct {
	Gets int64
	Hits int64
	MaxItemSize int
	CurrentSize int
}

type Cache interface {
	Set(key string, value interface{})
	Get(key string)(interface{}, bool)
	Delete(key string)
	Status()(*CacheStatus)
}

type AtomicInt int64

func (i *AtomicInt) Add(n int64){
	atomic.AddInt64((*int64)(i), n)
}

func (i *AtomicInt) Get() int64{
	return atomic.LoadInt64((*int64)(i))
}

type MemCache struct {
	mutex sync.Mutex
	maxItemSize int
	cacheList *list.List
	cache map[interface{}]*list.Element
	hits,gets AtomicInt
}

type entery struct {
	key, val interface{}
}

func NewMemcache(maxItemSize int)*MemCache{
	return &MemCache{
		maxItemSize: maxItemSize,
		cacheList: list.New(),
		cache: make(map[interface{}]*list.Element),
	}
}

func (m *MemCache) Status() *CacheStatus{
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return &CacheStatus{
		MaxItemSize: m.maxItemSize,
		CurrentSize: m.cacheList.Len(),
		Gets: m.gets.Get(),
		Hits: m.hits.Get(),
	}
}

func (m *MemCache) Get(key string) (interface{}, bool){
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.gets.Add(1)
	if ele, hit := m.cache[key]; hit{
		m.hits.Add(1)
		m.cacheList.MoveToFront(ele)
		return ele.Value.(*entery).val, true
	}

	return nil, false
}

func (m *MemCache) Set(key string, val interface{}){
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.cache == nil{
		m.cache =make(map[interface{}]*list.Element)
		m.cacheList = list.New()
	}

	if ele, ok := m.cache[key]; ok {
		m.cacheList.MoveToFront(ele)
		ele.Value.(*entery).val = val
		return
	}

	ele := m.cacheList.PushFront(&entery{key: key, val: val})
	m.cache[key] = ele
	if m.maxItemSize != 0 && m.cacheList.Len() > m.maxItemSize{
		m.RemoveOldest()
	}
}

func (m *MemCache)RemoveOldest(){
	if m.cache == nil{
		return
	}
	ele := m.cacheList.Back()
	if ele != nil{
		m.cacheList.Remove(ele)
		key := ele.Value.(*entery).key
		delete(m.cache, key)
	}
}

func (m *MemCache) Delete(key string){
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.cache == nil{
		return
	}

	if ele, ok := m.cache[key]; ok{
		m.cacheList.Remove(ele)
		key := ele.Value.(*entery).key
		delete(m.cache, key)
		return
	}
}

func main(){

}
