package login

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/helper"
	"github.com/songfei1983/go-api-server/internal/model"
	"github.com/songfei1983/go-api-server/internal/user"
	"net/http"
)

type Handler interface {
	Login(c echo.Context) error
	Logout(c echo.Context) error
	Register(c echo.Context) error
}

type handler struct {
	LoginUseCase UseCase
	UserUseCase  user.UseCase
}

func (h handler) Login(c echo.Context) error {
	data := &model.LoginUser{}
	if err := c.Bind(data); err != nil {
		return err
	}
	if err := validator.New().Struct(data); err != nil {
		return err
	}
	cc := helper.CustomContext{Context: c,}
	u, err := h.UserUseCase.GetByEmail(cc, data.Email)
	if err != nil {
		return err
	}
	if u.Password.Verify(data.Password) {
		t, err := helper.GenToken(*u)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return c.NoContent(http.StatusUnauthorized)
}

func (h handler) Logout(c echo.Context) error {
	panic("implement me")
}

func (h handler) Register(c echo.Context) error {
	data := &model.CreateUser{}
	if err := c.Bind(data); err != nil {
		return err
	}
	if err := validator.New().Struct(data); err != nil {
		return err
	}
	cc := helper.CustomContext{Context: c,}
	if err := h.UserUseCase.Create(cc, data); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func NewController(api *app.APP) error {
	api.Server.Logger.Info("Created login controller")

	userRepo := user.NewUserPersistence(api)
	userUseCase := user.NewUseCase(userRepo)
	loginUseCase := NewUseCase()
	h := NewHandler(loginUseCase, userUseCase)

	api.Server.POST("/login", h.Login)
	api.Server.GET("/logout", h.Login)
	api.Server.POST("/register", h.Register)
	return nil
}

func NewHandler(l UseCase, u user.UseCase) Handler {
	return &handler{
		LoginUseCase: l,
		UserUseCase:u,
	}
}
