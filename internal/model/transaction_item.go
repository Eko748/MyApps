package model

// TransactionItem represents an item in a transaction.
type TransactionItem struct {
	ID           int     `json:"id"`
	TransactionID int    `json:"transaction_id"`
	ProductID    int     `json:"product_id"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json:"price"`
}
