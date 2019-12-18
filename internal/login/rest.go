package login

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/helper"
	"github.com/songfei1983/go-api-server/internal/model"
	"net/http"
)

type Handler interface {
	Login(c echo.Context) error
	Logout(c echo.Context) error
	Register(c echo.Context) error
}

type handler struct {
	LoginUseCase UseCase
}

func (h handler) Login(c echo.Context) error {
	data := &model.DataLoginRequest{}
	if err := c.Bind(data); err != nil {
		return err
	}
	if err := validator.New().Struct(data); err != nil {
		return err
	}
	u := model.User{
		ID:        0,
		Name:      data.Username,
		Role:      "admin",
		Email:     data.Username,
		Password:  model.Password(data.Password),
		IsEnabled: true,
	}
	t, err := helper.GenToken(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func (h handler) Logout(c echo.Context) error {
	panic("implement me")
}

func (h handler) Register(c echo.Context) error {
	panic("implement me")
}

func NewController(api *app.APP) error {
	api.Server.Logger.Info("Created login controller")

	loginUseCase := NewUseCase()
	h := NewHandler(loginUseCase)

	api.Server.POST("/login", h.Login)
	api.Server.GET("/logout", h.Login)
	api.Server.POST("/register", h.Register)
	return nil
}

func NewHandler(login UseCase) Handler {
	return &handler{
		LoginUseCase: login,
	}
}
