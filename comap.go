package comap

import (
	"sync"
)

// Specify the number of the elements when CoMap is allocated.
var COUNT = 32

// ConcurrentMap is encapsulated into an array.
type CoMap []*ConcurrentMap

// The definition to thread-safe concurrent map.
type ConcurrentMap struct {
	concurrentMap map[string]interface{}
	sync.RWMutex
}

// Get the CoMap[key]'s value.
func (m CoMap) Get(key string) (interface{}, bool) {
	// Get elem
	elem := m.GetShard(key)
	elem.RLock()
	// Get ConcurrentMap from elem.
	val, ok := elem.concurrentMap[key]
	if !ok {
		panic("error occurred when executing elem.concurrentMap[key]")
	}
	elem.RUnlock()
	return val, ok
}

// Set the CoMap[key]'s value.
func (m CoMap) Set(key string, value interface{}) {
	// Get map elem.
	elem := m.GetShard(key)
	elem.Lock()
	elem.concurrentMap[key] = value
	elem.Unlock()
}

// Get the corresponding key's map: .
func (m CoMap) GetShard(key string) *ConcurrentMap {
	return m[uint(hash(key))%uint(COUNT)]
}

// TODO: improving...
func hash(key string) uint32 {
	hash := uint32(3124590231)
	const prime32 = uint32(19756321)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

// make a new CoMap
func New() CoMap {
	m := make(CoMap, COUNT)
	for i := 0; i < COUNT; i++ {
		m[i] = &ConcurrentMap{concurrentMap: make(map[string]interface{})}
	}
	return m
}
