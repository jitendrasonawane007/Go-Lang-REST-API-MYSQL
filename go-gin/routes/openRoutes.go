package routes

// package routes
// package routes

import (
	handler "go-gin/handlers"

	"github.com/gin-gonic/gin"
)

func OpenRoutes(o *gin.RouterGroup) {
	o.POST("/login", handler.Login())
	o.POST("/save", handler.SaveUserDetails())
}
