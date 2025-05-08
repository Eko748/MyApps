package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}

func JSON(w http.ResponseWriter, code int, message string, data interface{}, pagination interface{}) {
	// Prepare the API response
	res := APIResponse{
		Status:    code < 400,
		Message:   message,
		Code:      code,
		Data:      data,       // Directly assign the products data
		Pagination: pagination, // Add pagination if available
	}

	// Set the response header and encode the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}
