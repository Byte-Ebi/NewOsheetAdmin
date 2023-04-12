package admin

import (
	"NewOsheetAdmin/internal/authentication"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	v1 := r.Group("v1")
	v1.POST("/authentication")

	v1.Use(authentication.Middleware())

	v1.DELETE("/authentication")
}
