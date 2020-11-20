package main

import (
	middleware "go-gin/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	md := cors.DefaultConfig()
	md.AllowAllOrigins = true
	md.AllowHeaders = []string{"*"}
	md.AllowMethods = []string{"*"}
	md.ExposeHeaders = []string{"Authorization"}
	middleware.InitMiddleWare(router)
	s := &http.Server{
		Addr:    ":4000",
		Handler: router,
	}
	s.ListenAndServe()
}
