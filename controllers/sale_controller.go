package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/model"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/usecase"
)

type SaleController struct {
	usecase usecase.SaleUsecase
}

func NewSaleController(usecase usecase.SaleUsecase) SaleController {
	return SaleController{
		usecase: usecase,
	}
}

func (sc *SaleController) GetSales(ctx *gin.Context) {
	sales, err := sc.usecase.GetSales()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, sales)
}

func (sc *SaleController) CreateSale(ctx *gin.Context) {
	var sale model.Sale

	if err := ctx.ShouldBindJSON(&sale); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdSale, err := sc.usecase.CreateSale(sale)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar a venda."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Venda criada com sucesso",
		"product": createdSale,
	})
}
