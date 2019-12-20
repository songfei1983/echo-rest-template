package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strconv"
	"time"
)


type ZapLoggerConfig struct {
	Skipper middleware.Skipper
}

var DefaultLoggerConfig = ZapLoggerConfig{
	Skipper: middleware.DefaultSkipper,
}

func ZapLogger() echo.MiddlewareFunc {
	return ZapLoggerWithConfig(DefaultLoggerConfig)
}

func ZapLoggerWithConfig(config ZapLoggerConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			latency := stop.Sub(start)
			fields := defaultFields(c)
			appendFields := []zapcore.Field{
				zap.String("latency", strconv.FormatInt(int64(latency), 10)),
				zap.String("latency_human", latency.String()),
			}
			fields = append(fields, appendFields...)
			n := c.Response().Status
			switch {
			case n >= 500:
				global.Error("Server error", fields...)
			case n >= 400:
				global.Warn("Client error", fields...)
			case n >= 300:
				global.Info("Redirection", fields...)
			default:
				global.Info("Success", fields...)
			}
			return nil
		}
	}
}
