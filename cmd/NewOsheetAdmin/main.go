package main

import (
	"NewOsheetAdmin"
	"github.com/gin-gonic/gin"
)

func setupRoutes() *gin.Engine {
	r := NewOsheetAdmin.Router{}
	return r.Set()
}

func main() {
	r := setupRoutes()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
