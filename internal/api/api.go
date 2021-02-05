package api

import (
	"context"
	"sync"
	"time"

	_ "github.com/songfei1983/go-api-server/docs"
	"github.com/songfei1983/go-api-server/internal/server"
	"github.com/songfei1983/go-api-server/pkg/cache"
	"github.com/songfei1983/go-api-server/pkg/config"
)

var globalServer server.Server
var once sync.Once

const defaultTimeOut = 10 * time.Second

func Run(conf config.Config) {
	once.Do(func() {
		e := server.NewEcho(conf)
		p := cache.NewGoCache(conf)
		NewEchoHandler(e, p)
		globalServer = e
	})
	globalServer.Start()
}

func Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()
	return globalServer.Shutdown(ctx)
}
