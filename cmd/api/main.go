package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/songfei1983/go-api-server/handler/rest"
	"github.com/songfei1983/go-api-server/infra/persistence"
	"github.com/songfei1983/go-api-server/usecase"
)
import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	userRepository := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := rest.NewUserHandler(userUseCase)
	e.GET("/", userHandler.List)
	e.Logger.Fatal(e.Start(":1323"))
}

