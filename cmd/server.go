package cmd

import (
	"github.com/labstack/echo/v4"
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
		var options []func(*server.Server)
		if len(dataSource) > 0 {
			options = append(options, server.InitRepository(dataSource))
		}
		s, err := server.NewApp(options...)
		if err != nil {
			s.Mux.Logger.Fatal(err)
		}
		s.Mux.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				return h(&server.Context{Context: c, Server: s})
			}
		})
		routes(s)
		s.Mux.Logger.Fatal(s.Mux.Start(address))
	},
}

func routes(app *server.Server) {
	app.Mux.GET("/adunits", server.ActionFunc(adunit.List))
}
