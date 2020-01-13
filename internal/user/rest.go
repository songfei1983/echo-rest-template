package user

import (
	"context"
	"github.com/labstack/echo/v4/middleware"
	"github.com/songfei1983/go-api-server/internal/model"
	"github.com/songfei1983/go-api-server/pkg/helper"
	"github.com/songfei1983/go-api-server/pkg/logger"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/songfei1983/go-api-server/cmd/api/app"
)

func NewController(api *app.APP) error {
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
	api.Server.GET("/users/:id", userHandler.List, middlewares...)
	api.Server.POST("/users", userHandler.GetByID, middlewares...)
	api.Server.PUT("/users/:id", userHandler.Update2, middlewares...)
	api.Server.DELETE("/users/:id", userHandler.Delete, middlewares...)
	return nil
}

type handler struct {
	userUserCase UseCase
	Update2 echo.HandlerFunc
}

func newHandler(u UseCase) handler {
	h := handler{
		userUserCase: u,
	}
	h.Update2 = c(func(c *helper.CustomContext) error {
		return nil
	})
	return h
}

type callFunc func(c *helper.CustomContext) error
func c(h callFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(*helper.CustomContext))
	}
}

func (u handler) Update(c echo.Context) error {
	return nil
}

func (u handler) Delete(c echo.Context) error {
	c.Logger().Info("implement me")
	return nil
}

func (u handler) GetByID(c echo.Context) error {
	c.Logger().Info("implement me")
	return nil
}

func (u handler) List(c echo.Context) error {
	logger.Info(c, "list user")
	c.Set("AccountName", "Jason")
	ctx := context.WithValue(context.Background(), "ctx", c)
	res, err := u.userUserCase.GetAll(ctx)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, res)
}

func (u handler) Create(c echo.Context) error {
	type Request struct {
		Name     string `validate:"required"`
		Role     string `validate:"oneof=admin user"`
		Email    string `validate:"required,email"`
		Password string `validate:"required,gt=8,excludesall=;{}"`
	}
	req := new(Request)
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := validator.New().Struct(req); err != nil {
		c.Logger().Error(err)
		return err
	}

	m := new(model.CreateUser)
	m.Name = req.Name
	m.Email = req.Email
	m.Role = req.Role
	m.Password = model.Password(req.Password)

	logger.Info(c, "create user")
	c.Set("data", *m)
	ctx := helper.CustomContext{Context: c}
	if err := u.userUserCase.Create(ctx, m); err != nil {
		c.Logger().Error(err)
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusCreated)
}
