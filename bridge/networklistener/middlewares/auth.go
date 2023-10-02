package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secret string) func(c *gin.Context) {
	return func(c *gin.Context) {

		c.Next()
	}
}
