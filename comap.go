package comap

import (
	"sync"
	"strconv"
	"fmt"
)

var COUNT = 32

type ConcurrentMap []*Comap

// The definition to thread-safe concurrent map.
type Comap struct {
	comap map[string]interface{}
	sync.RWMutex
}

// TODO: figure out the reason for []*Comap

func (c ConcurrentMap) GetValue(key string) interface{} {
	val, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println(err)
	}
	cmap := c[val]
	defer cmap.RUnlock()
	return cmap.comap[key]
}

func (m ConcurrentMap) Set(key string, value interface{}) {
	// Get map shard.
	shard := m.GetShard(key)
	shard.Lock()
	shard.comap[key] = value
	shard.Unlock()
}

func (m ConcurrentMap) GetShard(key string) *Comap {
	return m[uint(fnv32(key))%uint(COUNT)]
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func New() ConcurrentMap {
	// TODO:figure out the reason for adding size in make()
	m := make(ConcurrentMap, COUNT)
	for i := 0; i < COUNT; i++ {
		m[i] = &Comap{comap: make(map[string]interface{})}
	}
	return m
}
