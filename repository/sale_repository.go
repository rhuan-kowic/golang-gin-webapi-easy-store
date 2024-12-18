package repository

import (
	"database/sql"
	"fmt"

	"github.com/rhuan-kowic/golang-gin-webapi-easy-store/model"
)

type SaleRepository struct {
	connection *sql.DB
}

func NewSaleRepository(connection *sql.DB) SaleRepository {
	return SaleRepository{
		connection: connection,
	}
}

func (sr *SaleRepository) GetSales() ([]model.Sale, error) {
	query := "SELECT id_sale, product_id, date_sale, total, quantity FROM sales"
	rows, err := sr.connection.Query(query)
	if err != nil {
		fmt.Println("Erro na consulta: ", err)
		return []model.Sale{}, err
	}
	var saleList []model.Sale
	var saleObj model.Sale

	for rows.Next() {
		err := rows.Scan(
			&saleObj.ID,
			&saleObj.ProductID,
			&saleObj.DateSale,
			&saleObj.Total,
			&saleObj.Quantity)

		if err != nil {
			fmt.Println(err)
			return []model.Sale{}, err
		}

		saleList = append(saleList, saleObj)
	}
	return saleList, nil
}
