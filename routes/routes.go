package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/controllers"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/usecase"
)

func SetupRouter(router *gin.Engine, productUsecase usecase.ProductUsecase, saleUsecase usecase.SaleUsecase) {

	// Controllers
	productController := controllers.NewProductController(productUsecase)
	saleController := controllers.NewSaleController(saleUsecase)

	// Rotas definidas
	router.GET("/products", productController.GetProducts)
	router.GET("/sales", saleController.GetSales)
}
