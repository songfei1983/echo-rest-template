package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	gormzap "github.com/wantedly/gorm-zap"

	api "github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/internal/infra"
	"github.com/songfei1983/go-api-server/internal/user/handler/rest"
	"github.com/songfei1983/go-api-server/logger"
)
import "github.com/labstack/echo/v4"

func main() {
	app := app()
	migrate(app)
	start(app)
}

func start(app *api.APP) {
	type Controller func(app *api.APP) error
	for _, handler := range []Controller{rest.NewUserController} {
		if err := handler(app); err != nil {
			log.Fatal(err)
		}
	}
	app.Server.Logger.Fatal(app.Server.Start(":1323"))
}

func app() *api.APP {
	app := api.NewApp()
	app.Server = server()
	app.DB = db(app)
	return app
}

func db(app *api.APP) *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.LogMode(true)
	db.SetLogger(gormzap.New(logger.New()))
	return db
}

func migrate(app *api.APP) {
	infra.Migrate(app.DB)
}

func server() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(logger.ZapLogger())
	e.Use(middleware.CORS())
	// e.Use(middleware.Recover())

	e.Logger.SetLevel(log.DEBUG)

	e.GET("/hc", func(c echo.Context) error {
		logger.Info(c, "hc")
		return c.String(http.StatusOK, "ok")
	})
	return e
}
