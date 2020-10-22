get-docs:
	go get -u github.com/swaggo/swag/cmd/swag

docs: get-docs
	swag init --dir cmd/api --parseDependency --output docs

build:
	go build -o bin/restapi cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./...

benchmark:
	go test -v ./... -bench . -benchmem -benchtime 1s -run ^$

all-test:
	go test -v ./... -bench . -benchmem -benchtime 1s

build-docker: build
	docker build . -t api-rest

run-docker: build-docker
	docker run -p 8888:8888 -e PORT=8888 api-rest

