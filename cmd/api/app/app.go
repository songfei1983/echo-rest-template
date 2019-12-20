package app

import (
	"github.com/songfei1983/go-api-server/pkg/helper"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/songfei1983/go-api-server/pkg/db"
	"github.com/songfei1983/go-api-server/pkg/logger"
	gormzap "github.com/wantedly/gorm-zap"
)

type APP struct {
	Config Config
	DB     *gorm.DB
	Server *echo.Echo
	Authorized *echo.Group
}

func New() *APP {
	app := new(APP)
	app.Server = app.server()
	app.DB = app.db()
	return app
}

type Config struct{}

func (app *APP) Start() {
	app.Server.Logger.Fatal(app.Server.Start(":1323"))
}

func (app *APP) db() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.LogMode(true)
	db.SetLogger(gormzap.New(logger.New()))
	return db
}

func (app *APP) Migrate() {
	db.Migrate(app.DB)
}

func (app *APP) server() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(logger.ZapLogger())
	e.Use(middleware.CORS())
	e.Use(middleware.JWTWithConfig(helper.DefaultJWTConfig))
	// e.Use(middleware.Recover())
	app.Authorized = e.Group("/api", helper.AuthenticationMiddleware)

	e.Logger.SetLevel(log.DEBUG)

	e.GET("/hc", func(c echo.Context) error {
		logger.Info(c, "hc")
		return c.String(http.StatusOK, "ok")
	})
	return e
}

func (app *APP) Close() {
	if app.DB != nil{
		if err := app.DB.Close(); err != nil {
			log.Error(err)
		}
	}
}
