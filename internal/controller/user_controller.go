package controller

import (
	"encoding/json"
	"fmt"
	"myapps/internal/model"
	"myapps/internal/config"
	"net/http"
)

// GetUsers - Mengambil daftar semua users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User

	// Query untuk mengambil semua data users
	rows, err := config.DB.Query(`
		SELECT id, email, password, otp, provider, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by 
		FROM users`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Memindai setiap row menjadi user
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.OTP, // Menyesuaikan dengan sql.NullString
			&user.Provider,
			&user.CreatedAt,
			&user.CreatedBy,
			&user.UpdatedAt,
			&user.UpdatedBy,
			&user.DeletedAt,
			&user.DeletedBy,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Menyusun respons JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUser - Membuat user baru
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	// Decode request body JSON ke struct User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Query untuk memasukkan data user ke dalam database
	query := `
		INSERT INTO users (email, password, otp, provider, created_at, created_by, updated_at, updated_by) 
		VALUES (?, ?, ?, ?, NOW(), ?, NOW(), ?)`
	_, err := config.DB.Exec(query, user.Email, user.Password, user.OTP, user.Provider, user.CreatedBy, user.UpdatedBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengirimkan respon sukses
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}
