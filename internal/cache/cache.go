package cache

import (
	"time"
)

type Instance struct {
	Cache InMemType
}

type InMemType map[string][]byte

func (i *Instance) Clear() error {
	i.Cache = make(InMemType)
	return nil
}

func (i *Instance) Delete(key []byte) error {
	k := string(key)
	if _, ok := i.Cache[k]; ok {
		delete(i.Cache, k)
		return nil
	} else {
		return nil
	}
}

func (i Instance) Get(key []byte) ([]byte, error) {
	if val, ok := i.Cache[string(key)]; ok {
		return val, nil
	} else {
		return nil, nil
	}
}

func (i *Instance) Set(key []byte, value []byte, ttl time.Duration) error {
	i.Cache[string(key)] = value
	return nil
}
