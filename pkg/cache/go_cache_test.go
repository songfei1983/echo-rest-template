package cache

import (
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

func BenchmarkGoCache_GET(b *testing.B) {
	c := &GoCache{
		client: cache.New(5*time.Minute, 10*time.Minute),
	}
	k, v := "test", "test value" //nolint:goconst
	c.PUT(k, v)                  //nolint:errcheck
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.GET(k) //nolint:errcheck
	}
}

func BenchmarkGoCache_PUT(b *testing.B) {
	c := &GoCache{
		client: cache.New(5*time.Minute, 10*time.Minute),
	}
	k, v := "test", "test value" //nolint:goconst
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.PUT(k, v) //nolint:errcheck
	}
}
