package adunit

import (
	"context"
	"net/http"

	"github.com/songfei1983/go-api-server/internal/server"
)

func Bind(path string, s *server.Server) {
	s.Mux.GET("/adunits", server.ActionFunc(list))
	s.Mux.POST("/adunits", server.ActionFunc(create))
}

func list(c *server.Context) error {
	ms, err := c.Server.Repo.QueryUser(context.Background())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func create(c *server.Context) error {
	var in struct {
		Name string
	}
	if err := c.Bind(&in); err != nil {
		return err
	}
	m, err := c.Server.Repo.CreateUser(context.Background(), in.Name)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, m)
}
