package cache

import (
	"time"
)

type Instance struct {
	cache inMemType
}

type inMemType map[string][]byte

func New() *Instance {
	return &Instance{
		cache: make(inMemType),
	}
}

func (i *Instance) Clear() error {
	i.cache = make(inMemType)
	return nil
}

func (i *Instance) Delete(key []byte) error {
	k := string(key)
	if _, ok := i.cache[k]; ok {
		delete(i.cache, k)
		return nil
	} else {
		return nil
	}
}

func (i *Instance) Get(key []byte) ([]byte, error) {
	if val, ok := i.cache[string(key)]; ok {
		return val, nil
	} else {
		return nil, nil
	}
}

func (i *Instance) Set(key []byte, value []byte, ttl time.Duration) error {
	i.cache[string(key)] = value
	return nil
}
