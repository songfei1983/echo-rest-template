package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/songfei1983/go-api-server/pkg/config"
	"github.com/songfei1983/go-api-server/pkg/logger"
)

type API struct {
	Config *config.Config
	DB     *gorm.DB
	Server *echo.Echo
}

func Open(conf *config.Config, options ...func(*API) error) (*API, error) {
	a := new(API)
	a.Config = conf
	for _, option := range options {
		err := option(a)
		if err != nil {
			panic(err)
		}
	}
	return a, nil
}

func DB() func(*API) error {
	return func(a *API) error {
		conn, err := gorm.Open(a.Config.DB.Driver, a.Config.DB.DBName)
		if err != nil {
			return err
		}
		conn.LogMode(a.Config.DB.LogMode)
		// conn.SetLogger(gormzap.New(logger.New()))
		conn.Debug()
		a.DB = conn
		return nil
	}
}

func Serve() func(*API) error {
	return func(a *API) error {
		e := echo.New()
		//e.Use(middleware.RequestID())
		//e.Use(logger.ZapLogger())
		// e.Use(middleware.CORS())
		// e.HTTPErrorHandler = customHTTPErrorHandler
		// e.Use(middleware.Recover())
		e.Use(middleware.Logger())
		e.Logger.SetLevel(log.DEBUG)
		e.GET("/hc", func(c echo.Context) error {
			logger.Info(c, "hc")
			return c.String(http.StatusOK, "ok")
		})
		a.Server = e
		return nil
	}
}

func (a *API) Start() {
	a.Server.Logger.Fatal(a.Server.Start(":1323"))
}

func (a *API) Close() {
	if a.DB != nil {
		if err := a.DB.Close(); err != nil {
			log.Error(err)
		}
	}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusBadRequest
	if he, ok := err.(validator.ValidationErrors); ok {
		c.JSON(code, he.Error())
	}
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		c.JSON(he.Code, he.Error())
	}
	c.Logger().Error(err)
}

func logLevel(s string) log.Lvl {
	lvm := map[string]log.Lvl{
		"debug": log.DEBUG,
		"info":  log.INFO,
		"error": log.ERROR,
		"warn":  log.WARN,
	}
	lv, ok := lvm[strings.ToLower(s)]
	if ok {
		return lv
	}
	return log.OFF
}
