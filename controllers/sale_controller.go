package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
