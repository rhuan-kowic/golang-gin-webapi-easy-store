package repository

import (
	"database/sql"
	"fmt"

	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id_product, name, description, price, stock_quantity, minimum_stock FROM products"

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println("Erro na consulta: ", err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Description,
			&productObj.Price,
			&productObj.StockQuantity,
			&productObj.MinimumStock)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, productObj)
	}
	rows.Close()
	return productList, nil
}
