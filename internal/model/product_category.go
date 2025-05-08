package model

// ProductCategory represents the many-to-many relationship between products and categories.
type ProductCategory struct {
	ProductID  int `json:"product_id"`
	CategoryID int `json:"category_id"`
}
