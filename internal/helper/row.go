package helper

import (
	"database/sql"
	"myapps/internal/model"
)

func ProductRows(rows *sql.Rows) ([]model.Product, error) {
	var products []model.Product

	for rows.Next() {
		var p model.Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Stock,
			&p.MainCategoryID,
			&p.CreatedAt,
			&p.CreatedBy,
			&p.UpdatedAt,
			&p.UpdatedBy,
			&p.DeletedAt,
			&p.DeletedBy,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
