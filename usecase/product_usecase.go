package usecase

import (
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/model"
	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}
