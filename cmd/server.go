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
	serverCmd.Flags().StringVar(&dataSource, "db", "", "file:ent?mode=memory&cache=shared&_fk=1")
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
			options = append(options, server.InitRepository("sqlite3", dataSource))
		}
		s, err := server.NewApp(options...)
		if err != nil {
			s.Mux.Logger.Fatal(err)
		}
		s.Mux.Use(s.WrapperContext)
		adunit.Bind("/adunits", s)
		s.Mux.Logger.Fatal(s.Mux.Start(address))
	},
}
