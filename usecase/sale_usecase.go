package usecase

import (
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/model"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/repository"
)

type SaleUsecase struct {
	repository repository.SaleRepository
}

func NewSaleUsecase(repo repository.SaleRepository) SaleUsecase {
	return SaleUsecase{
		repository: repo,
	}
}

func (su *SaleUsecase) GetSales() ([]model.Sale, error) {
	return su.repository.GetSales()
}

func (su *SaleUsecase) CreateSale(sale model.Sale) (model.Sale, error) {
	return su.repository.CreateSale(sale)
}
