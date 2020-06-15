package adunit

import (
	"net/http"

	"github.com/songfei1983/go-api-server/internal/server"
)

type Handler struct{}

func Bind(path string, s *server.Server) {
	s.Mux.GET("/adunits", server.ActionFunc(List))
}

func List(c *server.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "ok"})
}
