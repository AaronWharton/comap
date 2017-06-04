package comap

import (
	"sync"
)

// Specify the number of the elements when CoMap is allocated.
var COUNT = 32

type CoMap []*ConcurrentMap

// The definition to thread-safe concurrent map.
type ConcurrentMap struct {
	concurrentMap map[string]interface{}
	sync.RWMutex
}

func (m CoMap) Get(key string) (interface{}, bool) {
	// Get elem
	elem := m.GetShard(key)
	elem.RLock()
	// Get ConcurrentMap from elem.
	val, ok := elem.concurrentMap[key]
	elem.RUnlock()
	return val, ok
}

func (m CoMap) Set(key string, value interface{}) {
	// Get map elem.
	elem := m.GetShard(key)
	elem.Lock()
	elem.concurrentMap[key] = value
	elem.Unlock()
}

func (m CoMap) GetShard(key string) *ConcurrentMap {
	return m[uint(hash(key))%uint(COUNT)]
}

// TODO: more function ...
func hash(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func New() CoMap {
	m := make(CoMap, COUNT)
	for i := 0; i < COUNT; i++ {
		m[i] = &ConcurrentMap{concurrentMap: make(map[string]interface{})}
	}
	return m
}
