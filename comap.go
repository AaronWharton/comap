package comap

import (
	"sync"
)

// COUNT specify the number of the elements when CoMap is allocated.
const COUNT = 32

// CoMap encapsulates the ConcurrentMap into array.
type CoMap []*ConcurrentMap

// ConcurrentMap defines a thread-safe concurrent map.
type ConcurrentMap struct {
	concurrentMap map[string]interface{}
	sync.RWMutex
}

// Get gets the CoMap[key]'s value.
func (m CoMap) Get(key string) (interface{}) {
	// Get elem
	elem := m.GetShard(key)
	elem.RLock()
	// Get ConcurrentMap from elem.
	val, _ := elem.concurrentMap[key]
	elem.RUnlock()
	return val
}

// Set sets the CoMap[key]'s value.
func (m CoMap) Set(key string, value interface{}) {
	// Get map elem.
	elem := m.GetShard(key)
	elem.Lock()
	elem.concurrentMap[key] = value
	elem.Unlock()
}

// GetShard gets the corresponding key's map.
func (m CoMap) GetShard(key string) *ConcurrentMap {
	return m[uint(hash(key))%uint(COUNT)]
}

// hashing: bit shifting
func hash(key string) uint32 {
	var hash = uint32(len(key))
	// prime is a prime number to execute hashing bit shifting operation.
	const prime uint32 = 16777619
	for i := 0; i < len(key); i++ {
		hash = (hash << 4) ^ (hash >> 28) ^ uint32(key[i])
	}
	return hash % prime
}

// New creates a new CoMap with capacity COUNT.
func New() CoMap {
	m := make(CoMap, COUNT)
	for i := 0; i < COUNT; i++ {
		m[i] = &ConcurrentMap{concurrentMap: make(map[string]interface{})}
	}
	return m
}
