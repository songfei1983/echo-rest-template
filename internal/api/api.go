package api

import (
	"context"
	"time"

	"github.com/swaggo/echo-swagger"

	_ "github.com/songfei1983/go-api-server/docs"
	"github.com/songfei1983/go-api-server/internal/api/controllers"
	"github.com/songfei1983/go-api-server/internal/pkg/cache"
	"github.com/songfei1983/go-api-server/internal/pkg/config"
	"github.com/songfei1983/go-api-server/internal/pkg/server"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/
var globalServer *server.EchoServer

func Run(conf config.Config) {
	globalServer = server.NewEchoServer(conf)
	p := cache.NewGoCache(conf)
	c := controllers.NewEchoHandler(p)
	globalServer.Server().Debug = true
	globalServer.Server().Logger.SetHeader(`{"time":"${time_rfc3339}","level":"${level}","prefix":"${prefix}","file":"${long_file}","line":"${line}"}`)
	globalServer.Server().GET("/swagger/*", echoSwagger.WrapHandler)
	globalServer.Server().GET("/keys/:key", c.GetKey())
	globalServer.Server().PUT("/keys", c.AddKeyValue())
	globalServer.Start()
}

func Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return globalServer.Server().Shutdown(ctx)
}
