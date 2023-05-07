package NewOsheetAdmin

import (
	"NewOsheetAdmin/app/auth"
	"NewOsheetAdmin/app/companies"
	"NewOsheetAdmin/app/groups"
	"NewOsheetAdmin/app/talents"
	"NewOsheetAdmin/internal/cors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r Router) Set() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Middleware())

	router.GET("/hc", healthCheck)

	auth.Routes(router)
	talents.Routes(router)
	companies.Routes(router)
	groups.Routes(router)

	return router
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
