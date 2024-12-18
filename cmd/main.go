package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/controllers"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/db"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/repository"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/usecase"
	"log"
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
	SaleRepository := repository.NewSaleRepository(db)

	ProductUsecase := usecase.NewProductUseCase(ProductRepository)
	SaleUsecase := usecase.NewSaleUsecase(&SaleRepository)

	ProductController := controllers.NewProductController(ProductUsecase)
	SaleController := controllers.NewSaleController(SaleUsecase)

	server.GET("/products", ProductController.GetProducts)
	server.GET("/sales", SaleController.GetSales)
	server.Run(":3000")
}
