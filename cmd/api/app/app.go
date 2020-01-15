package app

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	gormzap "github.com/wantedly/gorm-zap"

	"github.com/songfei1983/go-api-server/pkg/config"
	"github.com/songfei1983/go-api-server/pkg/db"
	"github.com/songfei1983/go-api-server/pkg/logger"
)

type APP struct {
	Config     *config.Config
	DB         *gorm.DB
	Server     *echo.Echo
}

func New(c *config.Config) *APP {
	app := new(APP)
	app.Config = c
	app.Server = app.server()
	app.DB = app.db()
	return app
}

func (a *APP) Start() {
	a.Server.Logger.Fatal(a.Server.Start(":1323"))
}

func (a *APP) db() *gorm.DB {
	db, err := gorm.Open(a.Config.DB.Driver, a.Config.DB.DBName)
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.LogMode(true)
	db.SetLogger(gormzap.New(logger.New()))
	return db
}

func (a *APP) Migrate() {
	db.Migrate(a.DB)
}

func (a *APP) server() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(logger.ZapLogger())
	e.Use(middleware.CORS())
	e.HTTPErrorHandler = customHTTPErrorHandler
	// e.Use(middleware.Recover())
	e.Logger.SetLevel(logLevel(a.Config.LogLevel))
	e.GET("/hc", func(c echo.Context) error {
		logger.Info(c, "hc")
		return c.String(http.StatusOK, "ok")
	})
	return e
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

func (a *APP) Close() {
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
	}
	c.Logger().Error(err)
}

