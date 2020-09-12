package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/songfei1983/go-api-server/internal/persistent"
)

type EchoHandler struct {
	Persistent persistent.Cache
}

func NewEchoHandler(p persistent.Cache) *EchoHandler {
	return &EchoHandler{Persistent: p}
}

func (h *EchoHandler) Load() echo.HandlerFunc {
	return func(c echo.Context) error {
		var in struct {
			Key string `json:"key"`
		}
		if err := c.Bind(&in); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		v, err := h.Persistent.GET(in.Key)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		var out struct {
			Value interface{} `json:"value"`
		}
		out.Value = v
		return c.JSON(http.StatusOK, out)
	}
}

func (h *EchoHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var in struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}
		if err := c.Bind(&in); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err := h.Persistent.PUT(in.Key, in.Value); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.NoContent(http.StatusCreated)
	}
}
