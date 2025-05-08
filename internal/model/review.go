package model

import "time"

// Review represents a product review for a transaction item.
type Review struct {
	ID               int       `json:"id"`
	TransactionItemID int      `json:"transaction_item_id"`
	Rating           int       `json:"rating"` // Rating between 1 and 5
	Comment          string    `json:"comment,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	CreatedBy        int       `json:"created_by"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	UpdatedBy        int       `json:"updated_by,omitempty"`
	DeletedAt        time.Time `json:"deleted_at,omitempty"`
	DeletedBy        int       `json:"deleted_by,omitempty"`
}
