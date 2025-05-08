package controller

import (
    "encoding/json"
    "fmt"
    "myapps/internal/model"
    "myapps/internal/config"
    "net/http"
    "github.com/gorilla/mux"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
    var transactions []model.Transaction
    rows, err := config.DB.Query("SELECT * FROM transactions")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var transaction model.Transaction
        if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.TotalAmount, &transaction.Status, &transaction.CreatedAt); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        transactions = append(transactions, transaction)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(transactions)
}

func UpdateTransactionStatus(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    transactionID := vars["transaction_id"]

    var request struct {
        Status string `json:"status"`
    }
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := "UPDATE transactions SET status = ? WHERE id = ?"
    _, err := config.DB.Exec(query, request.Status, transactionID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Transaction status updated successfully")
}
