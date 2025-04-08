package middleware

import (
    "github.com/gin-gonic/gin"
    "log"
    "time"
)

func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()

        // Process the request
        c.Next()

        // Calculate duration
        end := time.Since(start)

        // Log the request details
        log.Printf(
            "%s %s %d %s %v",
            c.Request.Method,
            c.Request.URL.Path,
            c.Writer.Status(),
            c.ClientIP(),
            end,
        )
    }
}