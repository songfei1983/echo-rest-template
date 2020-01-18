package main

import (
	"github.com/labstack/gommon/log"
	"github.com/songfei1983/go-api-server/internal/login"
	"github.com/songfei1983/go-api-server/internal/server"
	"github.com/songfei1983/go-api-server/internal/user"
	"github.com/songfei1983/go-api-server/pkg/config"
)

func main() {
	// config
	conf := config.NewConfig(log.New(""))
	conf.InitFlag()
	apiConf := conf.ParseConfig()
	// api server
	api, err := server.Open(apiConf, server.DB(), server.Serve())
	if err != nil {
		panic(err)
	}
	defer api.Close()
	// handle
	registerHandler(api)
	api.Start()
}

func registerHandler(api *server.API) {
	type Controller func(api *server.API) error
	for _, handler := range []Controller{
		user.NewController,
		login.NewController,
	} {
		if err := handler(api); err != nil {
			log.Fatal(err)
		}
	}
}

