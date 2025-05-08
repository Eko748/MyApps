package model

import "time"

// Category represents a category in the system.
type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Description string  `json:"description,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UpdatedBy int       `json:"updated_by,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	DeletedBy int       `json:"deleted_by,omitempty"`
}
