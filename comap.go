package comap

import "sync"

// The definition to thread-safe concurrent map.
type Comap struct {
	comap	map[string]interface{}
	sync.RWMutex
}

// TODO: figure out the reason for []*Comap

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

func New() *Comap {
	// TODO:figure out the reason for adding size in make()
	return &Comap{comap:make(map[string]interface{})}
}