package main

import (
	"fmt"
	"github.com/songfei1983/go-api-server/internal/cli"
	"sync"
)

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

	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		fmt.Println("I don't want to coding anything")
	}()
	wg.Wait()
}
