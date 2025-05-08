package helper

import (
	"myapps/internal/model"
	"errors"
)

// ValidateProduct melakukan validasi terhadap data produk
func ValidateProduct(p model.Product) error {
	if p.Name == "" {
		return errors.New("product name is required")
	}
	if p.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}
	if p.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}
	if p.MainCategoryID == 0 {
		return errors.New("product category is required")
	}
	return nil
}
