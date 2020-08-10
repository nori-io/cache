package main

import (
	"context"

	cacheInstance "github.com/nori-io/cache-memory/internal/cache"
	"github.com/nori-io/common/v3/config"
	"github.com/nori-io/common/v3/logger"
	"github.com/nori-io/common/v3/meta"
	"github.com/nori-io/common/v3/plugin"
	"github.com/nori-io/interfaces/nori/cache"
)

type service struct {
	instance *cacheInstance.Instance
}

var (
	Plugin plugin.Plugin = &service{}
)

func (p *service) Init(ctx context.Context, config config.Config, log logger.FieldLogger) error {
	return nil
}

func (p *service) Instance() interface{} {
	return p.instance
}

func (p *service) Meta() meta.Meta {
	return &meta.Data{
		ID: meta.ID{
			ID:      "nori/cache/memory",
			Version: "1.0.0",
		},
		Author: meta.Author{
			Name: "Nori.io",
			URI:  "https://nori.io/",
		},
		Core: meta.Core{
			VersionConstraint: "=0.2.0",
		},
		Dependencies: []meta.Dependency{},
		Description: meta.Description{
			Name:        "Nori: Cache In Memory",
			Description: "",
		},
		Interface: cache.CacheInterface,
		License: []meta.License{
			{
				Title: "GPLv3",
				Type:  "GPLv3",
				URI:   "https://www.gnu.org/licenses/"},
		},
		Tags: []string{"cache", "memory"},
	}

}

func (p *service) Start(ctx context.Context, registry plugin.Registry) error {
	if p.instance == nil {
		instance := &cacheInstance.Instance{
			Cache: make(cacheInstance.InMemType),
		}
		p.instance = instance
	}
	return nil
}

func (p *service) Stop(ctx context.Context, registry plugin.Registry) error {
	p.instance = nil
	return nil
}
