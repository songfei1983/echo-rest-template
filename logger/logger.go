package logger

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var global *zap.Logger
var once sync.Once

func init() {
	New()
}
func New() *zap.Logger {
	once.Do(func(){
		l, _ := zap.NewProduction()
		global = l
	})
	return global
}

func defaultFields(c echo.Context) []zapcore.Field {
	req := c.Request()
	res := c.Response()
	id := req.Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXRequestID)
	}
	fields := []zapcore.Field{
		zap.String("prefix", "echo"),
		zap.String("time", time.Now().Format(time.RFC3339)),
		zap.String("id", id),
		zap.String("method", req.Method),
		zap.String("uri", req.RequestURI),
		zap.String("host", req.Host),
		zap.String("remote_ip", c.RealIP()),
		zap.String("user_agent", req.UserAgent()),
	}
	return fields
}

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

func loggerFields(c echo.Context) []zapcore.Field{
	_, file, line, _ := runtime.Caller(2)
	fields := defaultFields(c)
	appendFields := []zapcore.Field{
		zap.String("file", path.Base(file)),
		zap.Int("line", line),
	}
	fields = append(fields, appendFields...)
	return fields
}

func Info(c echo.Context, format string, args ...interface{}) {
	global.Info(fmt.Sprintf(format, args...), loggerFields(c)...)
}
