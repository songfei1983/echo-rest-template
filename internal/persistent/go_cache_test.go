package persistent

import (
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

func BenchmarkGoCache_GET(b *testing.B) {
	c := &GoCache{
		client: cache.New(5*time.Minute, 10*time.Minute),
	}
	k, v := "test", "test value"
	c.PUT(k, v)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.GET(k)
	}
}

func BenchmarkGoCache_PUT(b *testing.B) {
	c := &GoCache{
		client: cache.New(5*time.Minute, 10*time.Minute),
	}
	k, v := "test", "test value"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.PUT(k, v)
	}
}