package cmd

import (
	"github.com/songfei1983/go-api-server/internal/adunit"
	"github.com/songfei1983/go-api-server/internal/server"
	"github.com/spf13/cobra"
)

var (
	address    string
	dataSource string
)

var _ = initServerCmd()

func initServerCmd() struct{} {
	serverCmd.Flags().StringVar(&address, "address", ":1323", "host:port")
	serverCmd.Flags().StringVar(&dataSource, "db", "", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	// _ = serverCmd.MarkFlagRequired("db")

	rootCmd.AddCommand(serverCmd)
	return struct{}{}
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start server",
	Long:  `start a http(s) server`,
	Run: func(cmd *cobra.Command, args []string) {
		var options []func(*server.App)
		if len(dataSource) > 0 {
			options = append(options, server.DatabaseMySQL(dataSource))
		}
		s, err := server.NewApp(options...)
		if err != nil {
			s.Listener.Logger.Fatal(err)
		}
		routes(s)
		s.Listener.Logger.Fatal(s.Listener.Start(address))
	},
}

func routes(app *server.App) {
	adunitHandler := adunit.InitHandler()
	app.Listener.GET("/adunits", adunitHandler.List)
}
