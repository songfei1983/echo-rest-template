module github.com/songfei1983/go-api-server

// +heroku goVersion go1.15
// +heroku install ./cmd/...
go 1.15

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/go-openapi/spec v0.20.2 // indirect
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-redis/redis/v8 v8.5.0
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/labstack/echo/v4 v4.1.17
	github.com/labstack/gommon v0.3.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/nxadm/tail v1.4.5 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/spf13/cobra v1.1.1
	github.com/swaggo/echo-swagger v1.1.0
	github.com/swaggo/swag v1.7.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/mod v0.3.1-0.20200828183125-ce943fd02449 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
)
