package routes

import (
	"NewOsheetAdmin/internal/authentication"

	"github.com/gin-gonic/gin"
)

func RouteAdmins(r *gin.Engine) {

	admin := r.Group("/admin")

	v1 := admin.Group("v1")
	v1.POST("/authentication")

	v1.Use(authentication.Middleware())

	v1.DELETE("/authentication")

	talents := v1.Group("/talents")
	talents.GET("")
	talents.POST("")
	talents.GET("/:name")
	talents.PUT("/:name")
}
