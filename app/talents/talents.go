package talents

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	v1 := r.Group("/v1")
	talents := v1.Group("/talents")

	talents.GET("")
	talents.POST("")

	talents.GET("/:id")
	talents.PUT("/:id")
	talents.DELETE("/:id")
	talents.PUT("/:id/graduate")
}
