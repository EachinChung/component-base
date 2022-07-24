package middleware

import (
	"net/http"
	"regexp"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const maxAge = 12

// Cors add cors headers.
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           maxAge * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			reg := `^.*eachin-life\.com$`
			rgx := regexp.MustCompile(reg)
			return rgx.MatchString(origin)
		},
	})
}

// NoCache 是一个中间件函数，它附加响应头以防止客户端缓存 HTTP 响应。
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
}

// Secure 是一个中间件函数，它附加了安全和资源的响应头。
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")

	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}
}

// Store 注册中间件的商店。
//goland:noinspection GoUnusedGlobalVariable
var Store = map[string]gin.HandlerFunc{
	"Recovery":        Recovery(),
	"Secure":          Secure,
	"NoCache":         NoCache,
	"Cors":            Cors(),
	"Logger":          Logger(),
	"LoggerWithColor": LoggerWithColor(),
}
