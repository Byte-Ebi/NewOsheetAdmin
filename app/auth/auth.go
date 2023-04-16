package auth

import (
	"NewOsheetAdmin/internal/authentication"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestAuth struct {
	Account  string `json:"account" example:"nanami247"`
	Password string `json:"password" example:"password123"`
}

func Routes(r *gin.Engine) {
	v1 := r.Group("v1")
	v1.POST("/authentication", authHandler)

	v1.Use(authentication.Middleware())

	v1.DELETE("/authentication", revokeHandler)
}

func authHandler(c *gin.Context) {
	rqs := RequestAuth{}
	if err := c.BindJSON(&rqs); err != nil {
		fmt.Printf("auth error: %v \n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "request body must be JSON format.",
		})
		return
	}

	// TODO 讀取資料庫資料
	const account string = "nanami247"
	const password string = "password123"
	if rqs.Account != account || rqs.Password != password {
		fmt.Printf("login fail: %v \n", rqs.Account)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Please Check Account Or Password.",
		})
		return
	}

	jwtHelper := authentication.NewJTWHelper()
	token, err := jwtHelper.Generate(rqs.Account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}

func revokeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"access_token": "",
	})
}
