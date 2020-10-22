package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"

	"github.com/songfei1983/go-api-server/internal/pkg/config"
)

type Redis struct {
	client *redis.Client
	conf   config.Config
}

func NewRedis(conf config.Config) *Redis {
	opt, err := redis.ParseURL(conf.Persistent.Redis.DataSource)
	if err != nil {
		log.Fatal(err)
	}
	rdb := redis.NewClient(opt)
	return &Redis{client: rdb, conf: conf}
}

func (r *Redis) GET(k string) (v interface{}, err error) {
	ctx := context.Background()
	return r.client.Get(ctx, k).Result()
}

func (r *Redis) PUT(k string, v interface{}) error {
	ctx := context.Background()
	return r.client.Set(ctx, k, v, 0).Err()
}

func (r *Redis) DEL(k string) error {
	ctx := context.Background()
	return r.client.Del(ctx, k).Err()
}
