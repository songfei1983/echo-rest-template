package server

import (
	"context"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"

	"github.com/songfei1983/go-api-server/pkg/config"
	"github.com/songfei1983/go-api-server/pkg/validator"
)

type EchoServer struct {
	server *echo.Echo
	conf   config.Config
}

func NewEcho(conf config.Config) *EchoServer {
	e := echo.New()
	e.Debug = conf.Server.Debug
	e.Logger.SetHeader(`{"time":"${time_rfc3339}","level":"${level}","prefix":"${prefix}","file":"${long_file}","line":"${line}"}`)
	e.Validator = validator.New()
	if e.Debug {
		e.Logger.SetLevel(log.DEBUG)
		e.Logger.Debug("enable debug")
	}
	e.Use(middleware.Logger())
	return &EchoServer{
		server: e,
		conf:   conf,
	}
}

func (s *EchoServer) Start() {
	s.server.Logger.Info(s.server.Start(s.conf.Server.String()))
}

func (s *EchoServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *EchoServer) Server() *echo.Echo {
	return s.server
}
