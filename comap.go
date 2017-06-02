package comap

import "sync"

var SIZE int = 32

type value interface {}

// The definition to thread-safe concurrent map.
type Comap struct {
	comap	map[string]interface{}
	sync.RWMutex
}

// TODO: figure out the reason for []*Comap
type ConcurrentMap []*Comap

func (comap Comap) GetValue(key string) interface{} {
	comap.RLock()
	defer comap.RUnlock()
	return comap.comap[key]
}

func (comap Comap) SetValue(key string, value interface{}) {
	comap.Lock()
	defer comap.Unlock()
	comap.comap[key] = value
}

func New() ConcurrentMap {
	// TODO:figure out the reason for adding size in make()
	m := make(ConcurrentMap, SIZE)
	for i := 0; i < SIZE; i++ {
		m[i] = &Comap{comap:make(map[string]interface{})}
	}
	return m
}