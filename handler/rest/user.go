package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/songfei1983/go-api-server/usecase"
	"net/http"
)

type UserHandler interface {
	List(c echo.Context) error
}

type userHandler struct {
	userUserCase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUserCase:u,
	}
}

func (u userHandler) List(c echo.Context) error {
	res, err := u.userUserCase.GetAll(c.Request().Context())
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, res)
}

