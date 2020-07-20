package main

import (
	"github.com/songfei1983/go-api-server/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
