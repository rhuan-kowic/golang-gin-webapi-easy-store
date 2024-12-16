package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Easy Store",
		})
	})
	server.Run(":3000")
}
