package middleware

import (
	"fmt"
	"time"

	"github.com/eachinchung/log"
	"github.com/gin-gonic/gin"
)

type logFormatterParams struct {
	enableColor bool
	gin.LogFormatterParams
}

// logFormatter 是 Logger 中间件使用的默认日志格式函数。
var logFormatter = func(param logFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.enableColor {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}

	return fmt.Sprintf("%s%3d%s - [%s] \"%v %s%s%s %s\" %s",
		statusColor, param.StatusCode, resetColor,
		param.ClientIP,
		param.Latency,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

// Logger 实例一个 Logger 中间件。
func Logger() gin.HandlerFunc {
	return LoggerWithConfig(false)
}

// LoggerWithColor 实例一个带颜色的 Logger 中间件。
func LoggerWithColor() gin.HandlerFunc {
	return LoggerWithConfig(true)
}

// LoggerWithConfig 实例一个带有 config 的 Logger 中间件。
func LoggerWithConfig(enableColor bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		param := logFormatterParams{
			enableColor: enableColor,
			LogFormatterParams: gin.LogFormatterParams{
				Request: c.Request,
				Keys:    c.Keys,
			},
		}

		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path

		log.L(c).Info(logFormatter(param))
	}
}
