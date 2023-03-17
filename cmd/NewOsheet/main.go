package main

import (
	"NewOsheet/http/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSHandler())

	r.GET("/hc", healthCheck)
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
