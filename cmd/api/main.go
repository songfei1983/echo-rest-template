package main

import (
	"github.com/songfei1983/go-api-server/internal/login"
	"log"

	"github.com/songfei1983/go-api-server/cmd/api/app"
	"github.com/songfei1983/go-api-server/internal/user"
)

func main() {
	api := app.New()
	defer api.Close()
	api.Migrate()
	handler(api)
	api.Start()
}

func handler(api *app.APP) {
	type Controller func(api *app.APP) error
	for _, handler := range []Controller{
		user.NewController,
		login.NewController,
	} {
		if err := handler(api); err != nil {
			log.Fatal(err)
		}
	}
}

