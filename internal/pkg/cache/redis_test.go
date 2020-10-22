package cache

import (
	"testing"

	"github.com/go-redis/redis/v8"
)

func BenchmarkRedis_GET(b *testing.B) {
	c := Redis{client: redis.NewClient(&redis.Options{Addr: "localhost:6379"})}
	k, v := "test", "test value"
	c.PUT(k, v) //nolint:errcheck
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.GET(k) //nolint:errcheck
	}
}

func BenchmarkRedis_PUT(b *testing.B) {
	c := Redis{client: redis.NewClient(&redis.Options{Addr: "localhost:6379"})}
	k, v := "test", "test value"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.PUT(k, v) //nolint:errcheck
	}
}
