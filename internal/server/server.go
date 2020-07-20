package server

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"

	"github.com/songfei1983/go-api-server/ent"
	"github.com/songfei1983/go-api-server/ent/migrate"
	"github.com/songfei1983/go-api-server/internal/repository"
)

var app *Server
var once sync.Once

type Server struct {
	Mux     *echo.Echo
	Timeout time.Duration
	Repo    *repository.Repository
}

func NewApp(options ...func(*Server)) (*Server, error) {
	once.Do(func() {
		e := echo.New()
		e.Use(middleware.Logger())
		srv := Server{Mux: e}
		for _, option := range options {
			option(&srv)
		}
		app = &srv
	})
	return app, nil
}

func Timeout(t int) func(*Server) {
	return func(s *Server) {
		s.Timeout = time.Duration(t) * time.Second
	}
}

func InitRepository(driverName, dataSourceName string, options ...ent.Option) func(*Server) {
	return func(s *Server) {
		client, err := ent.Open(driverName, dataSourceName, options...)
		if err != nil {
			s.Mux.Logger.Fatal(err)
		}
		// Run migration.
		err = client.Schema.Create(
			context.Background(),
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
		)
		if err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
		s.Repo = repository.New(client)
	}
}
