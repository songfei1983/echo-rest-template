package login

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/internal/model"
	"net/http"
)

type handler struct {
	LoginUseCase UseCase
}

type LoginForm struct {
	Email    string `validate:"required,gt=0"`
	Password string `validate:"required,gt=0"`
}

func (h handler) Login(c echo.Context) error {
	data := new(LoginForm)
	if err := c.Bind(data); err != nil {
		return err
	}
	if err := validator.New().Struct(data); err != nil {
		return err
	}
	if token, err := h.LoginUseCase.Login(data); err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, map[string]string{"token": *token})
	}
}

func (h handler) Logout(c echo.Context) error {
	panic("implement me")
}

func (h handler) Register(c echo.Context) error {
	data := &model.User{}
	if err := c.Bind(data); err != nil {
		return err
	}
	if err := validator.New().Struct(data); err != nil {
		return err
	}
	if err := h.LoginUseCase.Register(data); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func NewController(api *app.APP) error {
	loginUseCase := NewUseCase(NewRepository(api), api.Server.Logger)
	h := newHandler(loginUseCase)

	api.Server.POST("/login", h.Login)
	api.Server.GET("/logout", h.Logout)
	api.Server.POST("/register", h.Register)
	return nil
}

func newHandler(l UseCase) *handler {
	return &handler{
		LoginUseCase: l,
	}
}
