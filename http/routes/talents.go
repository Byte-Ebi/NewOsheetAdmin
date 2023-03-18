package routes

import (
	"github.com/gin-gonic/gin"
)

func RouteChannels(r *gin.Engine) {
	v1 := r.Group("/v1")
	talents := v1.Group("/talents")

	talents.GET("")
	talents.GET("/:name")
}
