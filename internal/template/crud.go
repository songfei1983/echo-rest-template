package template

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type RequestSchema interface{}
type ResponseSchema interface{}
type RequestType string

const (
	RequestTypeList    RequestType = "List"
	RequestTypeGetByID RequestType = "GetByID"
	RequestTypeCreate  RequestType = "Create"
	RequestTypeUpdate  RequestType = "Update"
	RequestTypeDelete  RequestType = "Delete"
)

type Service interface {
	Create(c context.Context) error
	Update(c context.Context) error
	Delete(c context.Context) error
	Query(c context.Context) (ResponseSchema, error)
	RequestSchema(t RequestType) RequestSchema
}

type Query map[string]interface{}
type Pager struct {
	Limit   int
	Offset  int
	OrderBy string
}

type CRUD struct {
	Implement Service
	Group     *echo.Group
	Resource  string
}

func NewCRUD(g echo.Group, resource string, service Service) {
	crud := new(CRUD)
	crud.Group = g
	crud.Implement = service
	crud.Resource = resource
	crud.Group.GET("/"+crud.Resource, func(c echo.Context) error {
		data := crud.Implement.RequestSchema(RequestTypeList)
		if err := c.Bind(data); err != nil {
			return err
		}
		if err := validator.New().Struct(data); err != nil {
			return err
		}
		ctx := context.WithValue(context.Background(), "data", data)
		resp, err := crud.Implement.Query(ctx)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, resp)
	})
	crud.Group.GET("/"+crud.Resource+"/:id", func(c echo.Context) error {
		data := crud.Implement.RequestSchema(RequestTypeGetByID)
		if err := c.Bind(data); err != nil {
			return err
		}
		if err := validator.New().Struct(data); err != nil {
			return err
		}
		ctx := context.WithValue(context.Background(), "data", data)
		resp, err := crud.Implement.Query(ctx)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, resp)
	})
	crud.Group.POST("/"+crud.Resource, func(c echo.Context) error {
		ctx := context.WithValue(context.Background(), "c", c)
		data := crud.Implement.RequestSchema(RequestTypeCreate)
		if err := c.Bind(data); err != nil {
			return err
		}
		if err := validator.New().Struct(data); err != nil {
			return err
		}
		err := crud.Implement.Create(ctx)
		if err != nil {
			return err
		}
		return c.NoContent(http.StatusCreated)
	})
	crud.Group.PUT("/"+crud.Resource+"/:id", func(c echo.Context) error {
		data := crud.Implement.RequestSchema(RequestTypeUpdate)
		if err := c.Bind(data); err != nil {
			return err
		}
		if err := validator.New().Struct(data); err != nil {
			return err
		}
		ctx := context.WithValue(context.Background(), "data", data)
		err := crud.Implement.Update(ctx)
		if err != nil {
			return err
		}
		return c.NoContent(http.StatusNoContent)
	})
	crud.Group.DELETE("/"+crud.Resource+"/:id", func(c echo.Context) error {
		data := crud.Implement.RequestSchema(RequestTypeDelete)
		if err := c.Bind(data); err != nil {
			return err
		}
		if err := validator.New().Struct(data); err != nil {
			return err
		}
		ctx := context.WithValue(context.Background(), "data", data)
		err := crud.Implement.Delete(ctx)
		if err != nil {
			return err
		}
		return c.NoContent(http.StatusNoContent)
	})
}
