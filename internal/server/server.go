package server

import "context"

type Server interface {
	Start()
	Shutdown(ctx context.Context) error
}
