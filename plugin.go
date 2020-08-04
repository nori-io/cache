package cache

import (
	"context"
	cfg "github.com/nori-io/common/v3/config"
	"github.com/nori-io/common/v3/meta"
	noriPlugin "github.com/nori-io/common/v3/plugin"
	"github.com/nori-io/interfaces/cache"
	"plugin"
)

type service struct {
	instance *instance
}

type inMemType map[string][]byte

type instance struct {
	cache inMemType
}


var (
	Plugin plugin.Plugin = &service{}
)


func (p *service) Init(_ context.Context, configManager cfg.Manager) error {
	p.instance = &instance{}
	p.instance.cache =make(inMemType)
	return nil
}

func (p *service) Instance() interface{} {
	return p.instance
}

func (p service) Meta() meta.Meta {
	return &meta.Data{

		ID: meta.ID{
			ID:      "nori/cache/memory",
			Version: "1.0.0",
		},
		Author: meta.Author{
			Name: "Nori",
			URI:  "https://nori.io",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Dependencies: []meta.Dependency{},
		Description: meta.Description{
			Name: "Nori: Cache In Memory",
		},
		Interface:cache.CacheInterface ,
		License: meta.License{
			Title: "",
			Type:  "GPLv3",
			URI:   "https://www.gnu.org/licenses/"},
		Tags: []string{"cache", "memory"},
	}

}

func (p service) Start(ctx context.Context, registry noriPlugin.Registry) error {
	if p.instance == nil {
		instance := &instance{
			cache: make(inMemType),
		}

		p.instance = instance
	}
	return nil
}

func (p service) Stop(_ context.Context, _ noriPlugin.Registry) error {
	p.instance = nil
	return nil
}

