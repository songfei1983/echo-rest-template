package server

import (
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //import the database’s driver
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func InitRepository(dataSourceName string) func(*Server) {
	return func(s *Server) {
		db, err := gorm.Open("mysql", dataSourceName)
		if err != nil {
			s.Mux.Logger.Fatal(err)
		}
		s.Repo, err = repository.New(db)
		if err != nil {
			s.Mux.Logger.Fatal(err)
		}
	}
}
