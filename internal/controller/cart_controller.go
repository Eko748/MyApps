package controller

import (
    "encoding/json"
    "fmt"
    "myapps/internal/model"
    "myapps/internal/config"
    "net/http"
    "github.com/gorilla/mux"
)

func GetCartItems(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["user_id"]

    var cartItems []model.Cart
    query := "SELECT * FROM carts WHERE user_id = ?"
    rows, err := config.DB.Query(query, userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var cartItem model.Cart
        if err := rows.Scan(&cartItem.ID, &cartItem.UserID, &cartItem.ProductID, &cartItem.Quantity, &cartItem.CreatedAt); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        cartItems = append(cartItems, cartItem)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cartItems)
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
    var cartItem model.Cart
    if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := "INSERT INTO carts (user_id, product_id, quantity) VALUES (?, ?, ?)"
    _, err := config.DB.Exec(query, cartItem.UserID, cartItem.ProductID, cartItem.Quantity)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Product added to cart")
}
