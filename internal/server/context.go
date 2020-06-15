package server

import "github.com/labstack/echo/v4"

type Context struct {
	echo.Context
	Server *Server
}

type HandleFunc func(c *Context) error

func ActionFunc(h HandleFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(*Context))
	}
}
