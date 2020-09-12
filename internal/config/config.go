package config

import (
	"fmt"
	"time"
)

type Config struct {
	Server     Server
	Persistent Persistent
}

type Server struct {
	Hostname string
	Port     int
	Protocol string
}

func (s Server) String() string {
	return fmt.Sprintf("%s:%d", s.Hostname, s.Port)
}

type Persistent struct {
	GoCache GoCache
	Redis   Redis
}

type GoCache struct {
	DefaultExpiredTime time.Duration
	CleanupInterval    time.Duration
}

type Redis struct {
	DataSource string
}
