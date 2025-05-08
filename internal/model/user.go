package model

import (
	"time"
)

type User struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	OTP       *string    `json:"otp,omitempty"`       // pakai *string
	Provider  string     `json:"provider"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy int        `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *int       `json:"updated_by,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	DeletedBy *int       `json:"deleted_by,omitempty"`
}


