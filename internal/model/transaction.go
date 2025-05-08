package model

import "time"

// Transaction represents a user's transaction.
type Transaction struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"` // 'pending', 'paid', 'shipped', 'completed', 'cancelled'
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   int       `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	UpdatedBy   int       `json:"updated_by,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	DeletedBy   int       `json:"deleted_by,omitempty"`
}
