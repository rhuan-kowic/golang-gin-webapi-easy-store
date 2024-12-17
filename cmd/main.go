package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/db"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/repository"
)

func main() {
	server := gin.Default()
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	ProductRepository := repository.NewProductRepository(db)
	server.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Easy Store",
		})
	})

	server.GET("/products", func(ctx *gin.Context) {
		products, err := repository.GetProducts(&ProductRepository)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, products)
	})
	server.Run(":3000")
}
