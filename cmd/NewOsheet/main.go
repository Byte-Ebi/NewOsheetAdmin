package main

import (
	"NewOsheet/http/routes"
	"NewOsheet/internal/cors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Middleware())

	r.GET("/hc", healthCheck)

	routes.RouteAdmins(r)
	routes.RouteChannels(r)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
