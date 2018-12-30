package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func TestMid() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("testMiddleware")

		c.Next()
	}
}