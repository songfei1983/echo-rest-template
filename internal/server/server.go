package server

import (
	"github.com/songfei1983/go-api-server/internal/logger"
)

type Server interface {
	Start()
	Logger() logger.Logger
	Handle(method, path string, handler interface{}) error
}
