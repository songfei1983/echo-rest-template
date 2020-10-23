Echo REST API template
-----

[![Go Report Card](https://goreportcard.com/badge/github.com/songfei1983/go-api-server)](https://goreportcard.com/report/github.com/songfei1983/go-api-server)
[![Open Source Love](https://badges.frapsoft.com/os/mit/mit.svg?v=102)](https://github.com/ellerbrock/open-source-badge/)
![Go](https://github.com/songfei1983/go-api-server/workflows/Go/badge.svg)
![reviewdog](https://github.com/songfei1983/go-api-server/workflows/reviewdog/badge.svg)

## 1. Run with Docker

1. **Build**

```shell script
make build
docker build . -t api-rest
```

2. **Run**

```shell script
docker run -p 3000:3000 api-rest
```

3. **Test**

```shell script
go test -v ./test/...
```

_______

## 2. Generate Docs

```shell script
# Get swag
go get -u github.com/swaggo/swag/cmd/swag

# Generate docs
swag init --dir cmd/api --parseDependency --output docs
```

Run and go to **http://localhost:3000/docs/index.html**