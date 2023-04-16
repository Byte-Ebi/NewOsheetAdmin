package authentication

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "Forbidden",
			})
			c.Abort()
			fmt.Printf("Authorization is required")
			return
		}

		_, _, err := verifyToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Forbidden",
			})
			c.Abort()
			fmt.Printf("JWT Token Verify Fail: %v \n", err)
			return
		}
	}
}
