package middleware

import (
	"github.com/eachinchung/log"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const XRequestIDKey = "X-Request-ID"

// RequestID 是一个中间件，它将“X-Request-ID”注入每个请求的上下文和请求/响应标头中。
//goland:noinspection GoUnusedExportedFunction
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		rID := c.GetHeader(XRequestIDKey)
		if rID == "" {
			rID = uuid.NewV4().String()
			c.Request.Header.Set(XRequestIDKey, rID)
		}

		c.Set(log.KeyRequestID, rID)
		c.Writer.Header().Set(XRequestIDKey, rID)
		c.Next()
	}
}

// GetRequestIDFromContext returns 'RequestID' from the given context if present.
func GetRequestIDFromContext(c *gin.Context) string {
	if v, ok := c.Get(log.KeyRequestID); ok {
		if requestID, ok := v.(string); ok {
			return requestID
		}
	}

	return ""
}
