package cache

import (
	"time"
)



func (i *instance) Clear() error {
	i.cache = make(inMemType)
	return nil
}

func (i *instance) Delete(key []byte) error {
	k := string(key)
	if _, ok := i.cache[k]; ok {
		delete(i.cache, k)
		return nil
	} else {
		return nil
	}
}

func (i instance) Get(key []byte) ([]byte, error) {
	if val, ok := i.cache[string(key)]; ok {
		return val, nil
	} else {
		return nil, nil
	}
}

func (i *instance) Set(key []byte, value []byte, _ time.Duration) error {
	i.cache[string(key)] = value
	return nil
}