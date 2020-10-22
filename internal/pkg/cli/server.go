package cli

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/songfei1983/go-api-server/internal/api"
	"github.com/songfei1983/go-api-server/internal/pkg/config"
)

var (
	hostname string
	port     int
)

const (
	DefaultExpiredTime = 5 * time.Minute
	CleanupInterval    = 10 * time.Minute
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
			DefaultExpiredTime: DefaultExpiredTime,
			CleanupInterval:    CleanupInterval,
		}}}
		api.Run(conf)
	},
}
