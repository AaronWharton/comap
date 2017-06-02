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

func (c ConcurrentMap) SetValue(key string, value interface{}) {
	val, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println(err)
	}
	cmap := c[val]
	cmap.Lock()
	defer cmap.Unlock()
	cmap.comap[key] = value
}

func New() ConcurrentMap {
	// TODO:figure out the reason for adding size in make()
	m := make(ConcurrentMap, COUNT)
	for i := 0; i < COUNT; i++ {
		m[i] = &Comap{comap: make(map[string]interface{})}
	}
	return m
}
