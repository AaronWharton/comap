package comap

import "sync"

// The definition to thread-safe concurrent map.
type Comap struct {
	comap	map[string]interface{}
	sync.RWMutex
}