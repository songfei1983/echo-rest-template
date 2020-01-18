package user

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/songfei1983/go-api-server/pkg/helper"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/songfei1983/go-api-server/internal/server"
)

func NewController(api *server.API) error {
	api.Server.Logger.Info("New user handler")
	// inject
	userRepository := NewUserRepository(api)
	userUseCase := NewUseCase(userRepository)
	userHandler := newHandler(userUseCase)
	// router
	middlewares := []echo.MiddlewareFunc{
		middleware.JWTWithConfig(helper.DefaultJWTConfig), helper.AuthenticationMiddleware,
	}
	api.Server.GET("/users", userHandler.List, middlewares...)
	return nil
}

type handler struct {
	userUserCase UseCase
}

func newHandler(u UseCase) handler {
	h := handler{
		userUserCase: u,
	}
	return h
}

func (u handler) List(c echo.Context) error {
	res, err := u.userUserCase.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
