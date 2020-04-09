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

var app *App
var once sync.Once

type App struct {
	Listener *echo.Echo
	Timeout  time.Duration
	Repo     *repository.Repository
}

func NewApp(options ...func(*App)) (*App, error) {
	once.Do(func() {
		e := echo.New()
		e.Use(middleware.Logger())
		srv := App{Listener: e}
		for _, option := range options {
			option(&srv)
		}
		app = &srv
	})
	return app, nil
}

func Timeout(t int) func(*App) {
	return func(s *App) {
		s.Timeout = time.Duration(t) * time.Second
	}
}

func DatabaseMySQL(dataSourceName string) func(*App) {
	return func(s *App) {
		db, err := gorm.Open("mysql", dataSourceName)
		if err != nil {
			s.Listener.Logger.Fatal(err)
		}
		s.Repo, err = repository.New(db)
		if err != nil {
			s.Listener.Logger.Fatal(err)
		}
	}
}
