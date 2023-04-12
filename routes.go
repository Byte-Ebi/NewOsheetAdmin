package NewOsheetAdmin

import (
	"NewOsheetAdmin/app/admin"
	"NewOsheetAdmin/app/talents"
	"NewOsheetAdmin/internal/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
}

func (r Router) Set() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Middleware())

	router.GET("/hc", healthCheck)

	admin.Routes(router)
	talents.Routes(router)
	return router
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
