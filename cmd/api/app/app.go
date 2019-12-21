package app

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	gormzap "github.com/wantedly/gorm-zap"

	"github.com/songfei1983/go-api-server/pkg/config"
	"github.com/songfei1983/go-api-server/pkg/db"
	"github.com/songfei1983/go-api-server/pkg/helper"
	"github.com/songfei1983/go-api-server/pkg/logger"
)

type APP struct {
	Config *config.Config
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

func (a *APP) Start() {
	a.Server.Logger.Fatal(a.Server.Start(":1323"))
}

func (a *APP) db() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
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
	e.Use(middleware.JWTWithConfig(helper.DefaultJWTConfig))
	// e.Use(middleware.Recover())
	a.Authorized = e.Group("/api", helper.AuthenticationMiddleware)

	e.Logger.SetLevel(log.DEBUG)

	e.GET("/hc", func(c echo.Context) error {
		logger.Info(c, "hc")
		return c.String(http.StatusOK, "ok")
	})
	return e
}

func (a *APP) Close() {
	if a.DB != nil{
		if err := a.DB.Close(); err != nil {
			log.Error(err)
		}
	}
}
