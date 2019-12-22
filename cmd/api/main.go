package main

import (
	"github.com/labstack/gommon/log"
	"github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/internal/login"
	"github.com/songfei1983/go-api-server/internal/user"
	"github.com/songfei1983/go-api-server/pkg/config"
)

func main() {
	conf := config.NewConfig(log.New(""))
	conf.InitFlag()
	apiConf := conf.ParseConfig()
	api := app.New(apiConf)
	defer api.Close()
	api.Migrate()
	registerHandler(api)
	api.Start()
}

func registerHandler(api *app.APP) {
	type Controller func(api *app.APP) error
	for _, handler := range []Controller{
		user.NewController,
		login.NewController,
	} {
		if err := handler(api); err != nil {
			log.Fatal(err)
		}
	}
}

