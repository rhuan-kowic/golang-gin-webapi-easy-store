package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/db"
)

func main() {
	server := gin.Default()
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	server.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Easy Store",
		})
	})
	server.Run(":3000")
}
