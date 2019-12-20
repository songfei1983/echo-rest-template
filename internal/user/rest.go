package user

import (
	"context"
	"github.com/songfei1983/go-api-server/helper"
	"github.com/songfei1983/go-api-server/internal/model"
	"github.com/songfei1983/go-api-server/logger"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/songfei1983/go-api-server/cmd/api/app"
)

func NewController(api *app.APP) error {
	api.Server.Logger.Info("New user handler")
	// inject
	userRepository := NewUserPersistence(api)
	userUseCase := NewUseCase(userRepository)
	userHandler := NewHandler(userUseCase)
	// router
	api.Authorized.GET("/users", userHandler.List)
	api.Authorized.GET("/users/:id", userHandler.List)
	api.Authorized.POST("/users", userHandler.GetByID)
	api.Authorized.PUT("/users/:id", userHandler.Update)
	api.Authorized.DELETE("/users/:id", userHandler.Delete)
	return nil
}

type Handler interface {
	List(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetByID(c echo.Context) error
}

type handler struct {
	userUserCase UseCase
}

func NewHandler(u UseCase) Handler {
	return handler{
		userUserCase: u,
	}
}

func (u handler) Update(c echo.Context) error {
	c.Logger().Info("implement me")
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
	ctx := helper.CustomContext{Context:c}
	if err := u.userUserCase.Create(ctx, m); err != nil {
		c.Logger().Error(err)
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusCreated)
}
