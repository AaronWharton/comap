package comap

import (
	"sync"
	"fmt"
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
		fmt.Errorf("can not get val correctly when executing elem.concurrentMap[key]")
		// panic("error occurred when executing elem.concurrentMap[key]")
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

// Get the corresponding key's map.
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

// Create a new CoMap.
func New() CoMap {
	m := make(CoMap, COUNT)
	for i := 0; i < COUNT; i++ {
		m[i] = &ConcurrentMap{concurrentMap: make(map[string]interface{})}
	}
	return m
}
