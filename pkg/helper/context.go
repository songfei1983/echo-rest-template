package helper

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/songfei1983/go-api-server/internal/model"
	"net/http"
)

type Data interface{}

type CustomContext struct {
	echo.Context
	User *model.User
	Ctx  context.Context
	Data Data
}

func NewCustomContext(c echo.Context) *CustomContext {
	return &CustomContext{
		User:    new(model.User),
		Ctx:     nil,
		Data:    nil,
		Context: c,
	}
}

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		t, ok := c.Get("user").(*jwt.Token)
		if ok {
			if claims, ok := t.Claims.(*jwtCustomClaims); ok {
				cc := NewCustomContext(c)
				cc.User.ID = claims.ID
				cc.User.Name = claims.Name
				cc.User.Role = claims.Role
				cc.User.Email = claims.Email
				return next(cc)
			} else {
				c.Logger().Info("decode claims:", ok)
			}
		}
		c.Logger().Info("decode token :", ok)
		return c.NoContent(http.StatusUnauthorized)
	}
}
