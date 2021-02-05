get-docs:
	go get -u github.com/swaggo/swag/cmd/swag

docs: get-docs
	swag init --dir cmd --parseDependency --output docs

build:
	go build -o bin/api cmd/main.go

run:
	go run cmd/main.go

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

