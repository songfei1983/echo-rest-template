package rest

import (
	"context"
	"github.com/songfei1983/go-api-server/internal/user/domain/model"
	"github.com/songfei1983/go-api-server/logger"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	api "github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/internal/infra/persistence"
	"github.com/songfei1983/go-api-server/internal/user/usecase"
)

func NewUserController(app *api.APP) error {
	app.Server.Logger.Info("New user handler")
	userRepository := persistence.NewUserPersistence(app)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := NewUserHandler(userUseCase)
	app.Server.GET("/users", userHandler.List)
	app.Server.POST("/users", userHandler.Create)
	return nil
}

type UserHandler interface {
	List(c echo.Context) error
	Create(c echo.Context) error
}

type userHandler struct {
	userUserCase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return userHandler{
		userUserCase: u,
	}
}

func (u userHandler) List(c echo.Context) error {
	logger.Info(c, "list user")
	c.Set("AccountName", "Jason")
	ctx := context.WithValue(context.Background(), "ctx", c)
	res, err := u.userUserCase.GetAll(ctx)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, res)
}

func (u userHandler) Create(c echo.Context) error {
	type Request struct {
		Name string `validate:"required"`
		Age  uint   `validate:"required,gt=18"`
	}
	req := new(Request)
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := validator.New().Struct(req); err != nil {
		c.Logger().Error(err)
		return err
	}

	m := model.User{
		ID:   0,
		Name: req.Name,
		Age:  req.Age,
	}
	logger.Info(c, "create user")
	c.Set("data", m)
	ctx := context.WithValue(context.Background(), "ctx", c)
	if err := u.userUserCase.CreateUser(ctx); err != nil {
		c.Logger().Error(err)
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusCreated)
}
