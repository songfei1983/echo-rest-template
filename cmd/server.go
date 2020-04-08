package cmd

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var (
	address    string
	dataSource string
)

func init() {
	serverCmd.Flags().StringVar(&address, "address", ":1323", "host:port")
	serverCmd.Flags().StringVar(&dataSource, "db", "", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	_ = serverCmd.MarkFlagRequired("db")

	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start server",
	Long:  `start a http(s) server`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := NewServer(DatabaseMySQL(dataSource))
		if err != nil {
			s.Listener.Logger.Fatal(err)
		}
		s.Listener.Logger.Fatal(s.Listener.Start(address))
	},
}

type Server struct {
	Listener *echo.Echo
	Timeout  time.Duration
	DB       *gorm.DB
}

func NewServer(options ...func(*Server)) (*Server, error) {
	srv := Server{Listener: echo.New()}
	for _, option := range options {
		option(&srv)
	}
	return &srv, nil
}

func Timeout(t int) func(*Server) {
	return func(s *Server) {
		s.Timeout = time.Duration(t) * time.Second
	}
}

func DatabaseMySQL(dataSourceName string) func(*Server) {
	return func(s *Server) {
		db, err := gorm.Open("mysql", dataSourceName)
		if err != nil {
			s.Listener.Logger.Fatal(err)
		}
		s.DB = db
	}
}
