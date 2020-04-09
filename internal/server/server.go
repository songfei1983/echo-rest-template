package server

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //import the database’s driver
	"github.com/labstack/echo/v4"
)

type Server struct {
	Listener *echo.Echo
	Timeout  time.Duration
	DB       *gorm.DB
}

func NewServer(options ...func(*Server)) (*Server, error) {
	srv := Server{Listener: echo.New()}
	for _, option := range options {
		option(&srv)
	}
	return &srv, nil
}

func Timeout(t int) func(*Server) {
	return func(s *Server) {
		s.Timeout = time.Duration(t) * time.Second
	}
}

func DatabaseMySQL(dataSourceName string) func(*Server) {
	return func(s *Server) {
		db, err := gorm.Open("mysql", dataSourceName)
		if err != nil {
			s.Listener.Logger.Fatal(err)
		}
		s.DB = db
	}
}
