package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next() // Continue to the next handler
	}
}
