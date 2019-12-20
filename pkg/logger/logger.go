package logger

import (
	"fmt"
	"path"
	"runtime"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
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

func Debug(c echo.Context, format string, args ...interface{}) {
	global.Debug(fmt.Sprintf(format, args...), loggerFields(c)...)
}

func Info(c echo.Context, format string, args ...interface{}) {
	global.Info(fmt.Sprintf(format, args...), loggerFields(c)...)
}

func Warn(c echo.Context, format string, args ...interface{}) {
	global.Warn(fmt.Sprintf(format, args...), loggerFields(c)...)
}

func Error(c echo.Context, format string, args ...interface{}) {
	global.Error(fmt.Sprintf(format, args...), loggerFields(c)...)
}

func Fatal(c echo.Context, format string, args ...interface{}) {
	global.Fatal(fmt.Sprintf(format, args...), loggerFields(c)...)
}
