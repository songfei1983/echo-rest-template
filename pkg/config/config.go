package config

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Flag struct {
	Logger echo.Logger
	Path   *string
}

type Config struct {
	LogLevel string    `mapstructure:"log_level"`
	DB       Database  `mapstructure:"database"`
	API      APIServer `mapstructure:"api_server"`
}
type Database struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}
type APIServer struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func NewConfig(l echo.Logger) *Flag {
	return &Flag{
		Logger: l,
	}
}

func (c *Flag) InitFlag() {
	c.Path = pflag.String("config", "./configs/dev.yaml", "cmd --config /path/to/your/config")
}

func (c *Flag) ParseConfig() *Config {
	pflag.Parse()
	if c.Path == nil {
		c.Logger.Fatal("not set config")
	}
	c.Logger.Info(*c.Path)
	viper.SetConfigFile(*c.Path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		c.Logger.Fatal(err)
	}
	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		c.Logger.Fatal("unable to decode into config struct,", err)
	}
	c.Logger.Info(conf)
	return conf
}
