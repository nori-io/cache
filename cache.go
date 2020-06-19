package cache

import (
	"time"

	"github.com/nori-io/nori-common/v2/errors"
	"github.com/nori-io/nori-common/v2/meta"
	"github.com/nori-io/nori-common/v2/plugin"
)

const (
	CacheInterface meta.Interface = "nori/cache/Cache@1.0.0"
)

type Cache interface {
	Clear() error
	Delete(key []byte) error
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte, _ time.Duration) error
}

func GetCache(r plugin.Registry) (Cache, error) {
	instance, err := r.Interface(CacheInterface)
	if err != nil {
		return nil, err
	}
	i, ok := instance.(Cache)
	if !ok {
		return nil, errors.InterfaceAssertError{
			Interface: CacheInterface,
		}
	}
	return i, nil
}
