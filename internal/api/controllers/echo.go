package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/songfei1983/go-api-server/internal/pkg/cache"
	"github.com/songfei1983/go-api-server/internal/pkg/model"
	"github.com/songfei1983/go-api-server/internal/pkg/server"
	_ "github.com/songfei1983/go-api-server/pkg/errors"
)

type EchoHandler struct {
	Server *server.EchoServer
	Cache  cache.Cache
}

func NewEchoHandler(s *server.EchoServer, p cache.Cache) {
	h := &EchoHandler{Cache: p}
	s.Server().GET("/swagger/*", echoSwagger.WrapHandler)
	s.Server().GET("/keys/:key", h.GetKeyValue())
	s.Server().PUT("/keys", h.AddKeyValue())
}

// GetKeyValue godoc
// @Summary GetKeyValue a key
// @Description get string by key
// @ID get-value-by-key
// @Accept  json
// @Produce  json
// @Param key path int true "key"
// @Success 200 {object} model.KeyValue
// @Failure 400 {object} errors.HTTPError
// @Router /keys/{key} [get]
func (h *EchoHandler) GetKeyValue() echo.HandlerFunc {
	return func(c echo.Context) error {
		var in model.KeyValue
		if err := c.Bind(&in); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		c.Logger().Info(in)
		v, err := h.Cache.GET(in.Key)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		in.Value = v
		return c.JSON(http.StatusOK, in)
	}
}

// AddKeyValue godoc
// @Summary create key value
// @Description create an pair of key value
// @Accept  json
// @Produce  json
// @Success 204
// @Failure 400 {object} errors.HTTPError
// @Router /keys [put]
func (h *EchoHandler) AddKeyValue() echo.HandlerFunc {
	return func(c echo.Context) error {
		var in struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}
		if err := c.Bind(&in); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err := h.Cache.PUT(in.Key, in.Value); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.NoContent(http.StatusCreated)
	}
}
