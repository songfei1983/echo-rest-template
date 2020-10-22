package server

import (
	"github.com/labstack/echo/v4"

	"github.com/songfei1983/go-api-server/internal/pkg/config"
)

type EchoServer struct {
	server *echo.Echo
	conf   config.Config
}

func NewEchoServer(conf config.Config) *EchoServer {
	return &EchoServer{
		server: echo.New(),
		conf:   conf,
	}
}

func (s *EchoServer) Start() {
	s.server.Logger.Info(s.server.Start(s.conf.Server.String()))
}

func (s *EchoServer) Server() *echo.Echo {
	return s.server
}
