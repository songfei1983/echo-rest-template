FROM golang:1.15-alpine AS build_base

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN apk add --no-cache git  git gcc g++

# Set the Current Working Directory inside the container
WORKDIR /src

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/app ./cmd/main.go

# Start fresh from a smaller image
FROM alpine:3.12
RUN apk add ca-certificates

COPY --from=build_base /src/out/app /app/api

# Run the binary program produced by `go install`
CMD /app/api server --host 0.0.0.0 --port $PORT