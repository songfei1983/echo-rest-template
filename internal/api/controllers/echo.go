package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/songfei1983/go-api-server/internal/pkg/cache"
	"github.com/songfei1983/go-api-server/internal/pkg/model"
)

type EchoHandler struct {
	Persistent cache.Cache
}

func NewEchoHandler(p cache.Cache) *EchoHandler {
	return &EchoHandler{Persistent: p}
}

// GetKey godoc
// @Summary GetKey a key
// @Description get string by key
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "key"
// @Success 200 {object} model.KeyValue
// @Failure 400
// @Router /keys/{id} [get]
func (h *EchoHandler) GetKey() echo.HandlerFunc {
	return func(c echo.Context) error {
		var in model.KeyValue
		if err := c.Bind(&in); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		c.Logger().Info(in)
		v, err := h.Persistent.GET(in.Key)
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
// @Failure 400
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
		if err := h.Persistent.PUT(in.Key, in.Value); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.NoContent(http.StatusCreated)
	}
}
