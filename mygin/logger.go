package mygin

import (
	"log"
	"time"
)

// Logger 日志中间件
func Logger() HandlerFunc {
	return func(c *Context) {
		start := time.Now()
		path := c.Req.URL.Path
		statusCode := c.StatusCode

		c.Next()

		log.Printf("[%d] %s in %v", statusCode, path, time.Since(start))
	}
}
