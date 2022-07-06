package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func authorizationFailed(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, message)
}

func TokenAuthMiddleware() gin.HandlerFunc {
	environmentToken := os.Getenv("API_TOKEN")
	if environmentToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}
	return func(c *gin.Context) {
		token := c.Request.Header.Get("api_token")
		if token == "" {
			authorizationFailed(c, 400, "API token required")
			return
		}
		if token != token {
			authorizationFailed(c, 401, "Invalid API token")
			return
		}
		c.Next()
	}
}
