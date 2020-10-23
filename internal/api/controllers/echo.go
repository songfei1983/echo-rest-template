package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/songfei1983/go-api-server/internal/pkg/cache"
	"github.com/songfei1983/go-api-server/internal/pkg/model"
	"github.com/songfei1983/go-api-server/internal/pkg/server"
	"github.com/songfei1983/go-api-server/pkg/errors"
	_ "github.com/songfei1983/go-api-server/pkg/errors"
)

type EchoHandler struct {
	Server *server.EchoServer
	Cache  cache.Cache
}

func NewEchoHandler(s *server.EchoServer, p cache.Cache) {
	h := &EchoHandler{Cache: p}
	s.Server().GET("/", h.Top)
	s.Server().GET("/swagger/*", echoSwagger.WrapHandler)
	s.Server().GET("/keys/:key", h.GetKeyValue)
	s.Server().PUT("/keys", h.AddKeyValue)
}

func (h *EchoHandler) Top(c echo.Context) error {
	return c.String(http.StatusOK, "Echo REST API Template")
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
func (h *EchoHandler) GetKeyValue(c echo.Context) error {
	var in model.KeyValue
	if err := h.bindValidate(c, &in); err != nil {
		return c.JSON(err.Code, err)
	}
	c.Logger().Info(in)
	v, err := h.Cache.GET(in.Key)
	if err != nil {
		e := errors.NewHTTPError(http.StatusBadRequest, err)
		return c.JSON(e.Code, e)
	}
	in.Value = v
	return c.JSON(http.StatusOK, in)
}

// AddKeyValue godoc
// @Summary create key value
// @Description create an pair of key value
// @Accept  json
// @Produce  json
// @Success 204
// @Failure 400 {object} errors.HTTPError
// @Router /keys [put]
func (h *EchoHandler) AddKeyValue(c echo.Context) error {
	var in struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if err := h.bindValidate(c, &in); err != nil {
		return c.JSON(err.Code, err)
	}
	if err := h.Cache.PUT(in.Key, in.Value); err != nil {
		e := errors.NewHTTPError(http.StatusBadRequest, err)
		return c.JSON(e.Code, e)
	}
	return c.NoContent(http.StatusCreated)
}

func (h *EchoHandler) bindValidate(c echo.Context, in interface{}) *errors.HTTPError {
	if err := c.Bind(&in); err != nil {
		return errors.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(&in); err != nil {
		return errors.NewHTTPError(http.StatusBadRequest, err)
	}
	return nil
}
