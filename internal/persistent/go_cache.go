package persistent

import (
	"errors"

	"github.com/patrickmn/go-cache"

	"github.com/songfei1983/go-api-server/internal/config"
)

type GoCache struct {
	client *cache.Cache
	conf   config.Config
}

func NewGoCache(conf config.Config) *GoCache {
	return &GoCache{
		client: cache.New(conf.Persistent.GoCache.DefaultExpiredTime, conf.Persistent.GoCache.DefaultExpiredTime),
		conf:   conf,
	}
}

func (c *GoCache) GET(k string) (v interface{}, err error) {
	value, ok := c.client.Get(k)
	if !ok {
		return nil, errors.New("not found")
	}
	return value, nil
}

func (c *GoCache) PUT(k string, v interface{}) error {
	c.client.Set(k, v, cache.DefaultExpiration)
	return nil
}

func (c *GoCache) DEL(k string) error {
	c.client.Delete(k)
	return nil
}
