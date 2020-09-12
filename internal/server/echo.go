package server

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"

	"github.com/songfei1983/go-api-server/internal/config"
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
	s.server.Logger.Fatal(s.server.Start(s.conf.Server.String()))
}

func (s *EchoServer) Logger() echo.Logger {
	return s.server.Logger
}

func (s *EchoServer) Handle(method, path string, handler interface{}) error {
	ehc, ok := handler.(echo.HandlerFunc)
	if !ok {
		s.Logger().Error(reflect.TypeOf(handler))
		return errors.New("unknown handler type")
	}
	switch method {
	case http.MethodGet:
		s.server.GET(path, ehc)
	case http.MethodPut:
		s.server.PUT(path, ehc)
	case http.MethodPost:
		s.server.POST(path, ehc)
	case http.MethodDelete:
		s.server.DELETE(path, ehc)
	case http.MethodPatch:
		s.server.PATCH(path, ehc)
	default:
		return errors.New("unknown method")
	}
	return nil
}
