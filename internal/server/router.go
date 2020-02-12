package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewRouter(g *echo.Group, actions ...Action) {
	for _, a := range actions {
		fmt.Println(a.GetMethod(), a.GetPath())
		switch a.GetMethod() {
		case http.MethodGet:
			g.GET(a.GetPath(), a.Execute)
		case http.MethodPost:
			g.POST(a.GetPath(), a.Execute)
		case http.MethodPut:
			g.PUT(a.GetPath(), a.Execute)
		case http.MethodDelete:
			g.DELETE(a.GetPath(), a.Execute)
		}
	}
}

