package model

// Product represents a product in the system.
type Product struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Description    *string  `json:"description"`
	Price          float64 `json:"price"`
	Stock          int     `json:"stock"`
	MainCategoryID int     `json:"main_category_id"`
	CreatedAt      string  `json:"created_at"`
	CreatedBy      int     `json:"created_by"`
	UpdatedAt      *string `json:"updated_at,omitempty"`
	UpdatedBy      *int64  `json:"updated_by,omitempty"`
	DeletedAt      *string `json:"deleted_at,omitempty"`
	DeletedBy      *int64  `json:"deleted_by,omitempty"`
}
