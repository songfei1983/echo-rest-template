package main

import (
	"log"

	api "github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/internal/user/handler/rest"
)

func main() {
	app := api.New()
	defer app.Close()
	app.Migrate()
	handler(app)
	app.Start()
}

func handler(app *api.APP) {
	type Controller func(app *api.APP) error
	for _, handler := range []Controller{rest.NewUserController} {
		if err := handler(app); err != nil {
			log.Fatal(err)
		}
	}
}

