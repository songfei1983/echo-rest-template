package server

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"reflect"
)

type Response struct {
	StatusCode int
	Data       interface{}
}
type Service interface {
	Execute(c context.Context) (*Response, error)
	RequestSchema() interface{}
}

type Action interface {
	Execute(c echo.Context) error
	GetPath() string
	GetMethod() string
}

type ActionHandle struct {
	Method    string
	Path      string
	Implement Service
}

func (h *ActionHandle) GetPath() string {
	return h.Path
}

func (h *ActionHandle) GetMethod() string {
	return h.Method
}

func (h *ActionHandle) Execute(c echo.Context) error {
	data := h.Implement.RequestSchema()
	c.Logger().Info(reflect.TypeOf(data))
	if err := c.Bind(data); err != nil {
		return err
	}
	if err := validator.New().Struct(data); err != nil {
		return err
	}
	c.Logger().Info(data)
	ctx := context.WithValue(context.Background(), "data", data)
	resp, err := h.Implement.Execute(ctx)
	c.Logger().Info(resp, err)
	if err != nil {
		return err
	}
	if resp.Data == nil {
		return c.NoContent(resp.StatusCode)
	}
	return c.JSON(resp.StatusCode, resp.Data)
}

func NewAction(m, p string, s Service) Action {
	return &ActionHandle{Implement: s, Path: p, Method: m}
}
