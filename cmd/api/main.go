package main

import "github.com/songfei1983/go-api-server/internal/pkg/cli"

func main() {
	if err := cli.Execute(); err != nil {
		panic(err)
	}
}
