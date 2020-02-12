package main

import (
	"github.com/labstack/gommon/log"
	"github.com/songfei1983/go-api-server/internal/server"
	"github.com/songfei1983/go-api-server/internal/user"
	"github.com/songfei1983/go-api-server/pkg/config"
	"net/http"
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
	InitRouter(api)
	api.Start()
}

func InitRouter(api *server.API) {
	g := api.Server.Group("/api")
	actions := []server.Action{
		server.NewAction(http.MethodGet, "/users", user.NewList(api)),
		server.NewAction(http.MethodPost, "/login", user.NewLogin(api)),
		server.NewAction(http.MethodPost, "/register", user.NewRegister(api)),
	}
	server.NewRouter(g, actions...)
}
