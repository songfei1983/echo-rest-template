package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"

	"github.com/songfei1983/go-api-server/internal/config"
	"github.com/songfei1983/go-api-server/internal/handler"
	"github.com/songfei1983/go-api-server/internal/persistent"
	"github.com/songfei1983/go-api-server/internal/server"
)

var (
	hostname string
	port     int
)

var _ = initServerCmd()

func initServerCmd() struct{} {
	serverCmd.Flags().StringVar(&hostname, "hostname", "localhost", "hostname")
	serverCmd.Flags().IntVar(&port, "port", 8888, "port")

	rootCmd.AddCommand(serverCmd)
	return struct{}{}
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start server",
	Long:  `start a http(s) server`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.Config{Server: config.Server{
			Hostname: hostname,
			Port:     port,
			Protocol: "http",
		}, Persistent: config.Persistent{GoCache: config.GoCache{
			DefaultExpiredTime: 5 * time.Minute,
			CleanupInterval:    10 * time.Minute,
		}}}
		s := server.NewEchoServer(conf)
		p := persistent.NewGoCache(conf)
		c := handler.NewEchoHandler(p)
		if err := s.Handle(http.MethodGet, "/key", c.Load()); err != nil {
			log.Fatal(err)
		}
		if err := s.Handle(http.MethodPut, "/key", c.Create()); err != nil {
			log.Fatal(err)
		}
		s.Start()
	},
}
