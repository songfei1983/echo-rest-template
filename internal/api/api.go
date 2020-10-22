package api

import (
	"github.com/labstack/gommon/log"

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
func Run(conf config.Config) {
	s := server.NewEchoServer(conf)
	p := cache.NewGoCache(conf)
	c := controllers.NewEchoHandler(p)
	s.Server().Logger.SetLevel(log.DEBUG)
	s.Server().GET("/keys/:key", c.GetKey())
	s.Server().PUT("/keys", c.AddKeyValue())
	s.Start()
}
