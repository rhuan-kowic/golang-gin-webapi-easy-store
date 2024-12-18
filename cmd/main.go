package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/db"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/repository"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/routes"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/usecase"
)

func main() {
	// Iniciando o servidor
	server := gin.Default()

	// Conectando ao DB
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	// Criando os Repositorys
	productRepository := repository.NewProductRepository(db)
	saleRepository := repository.NewSaleRepository(db)

	// Criando os usecases
	productUsecase := usecase.NewProductUseCase(productRepository)
	saleUsecase := usecase.NewSaleUsecase(saleRepository)

	// Configura as rotas
	routes.SetupRouter(server, productUsecase, saleUsecase)

	// Inicia o servidor na porta 3000
	server.Run(":3000")
}
