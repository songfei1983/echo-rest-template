package api

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)
type APP struct {
	Config Config
	DB *gorm.DB
	Server *echo.Echo
}
func NewApp() *APP {
	return &APP{}
}

type Config struct {}

