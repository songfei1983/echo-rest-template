package main

import "github.com/songfei1983/go-api-server/internal/pkg/cli"

// @Golang API REST
// @version 1.0
// @description API REST in Golang with Echo Framework

// @contact.name fei song

// @license.name MIT
// @license.url https://raw.githubusercontent.com/songfei1983/go-api-server/master/LICENSE

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := cli.Execute(); err != nil {
		panic(err)
	}
}
