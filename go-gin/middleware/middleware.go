package middleware

import (
	"fmt"
	"go-gin/routes"

	"github.com/gin-gonic/gin"
)

/********************************************************/

func InitMiddleWare(g *gin.Engine) {
	OR := g.Group("/open")
	// OR.Use(OpenNetworkMiddleware())
	OR.Use(JSONMiddleware())
	routes.OpenRoutes(OR)
}

/********************************************************/
func OpenNetworkMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

/********************************************************/
func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(`check this from middleware`)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}
