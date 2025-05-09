package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"myapps/internal/response"
)

type OpenSourceController struct{}

// NewOpenSourceController creates a new instance of OpenSourceController
func NewOpenSourceController() *OpenSourceController {
	return &OpenSourceController{}
}

// GetWikipediaSuggestion calls the Wikipedia API and returns search suggestions
func (c *OpenSourceController) GetWikipediaSuggestion(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	if keyword == "" {
		response.JSON(w, http.StatusBadRequest, "Keyword is required", nil, nil)
		return
	}

	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=opensearch&format=json&search=%s", keyword)

	resp, err := http.Get(url)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Failed to fetch Wikipedia data", nil, nil)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Failed to read response", nil, nil)
		return
	}

	var result []interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		response.JSON(w, http.StatusInternalServerError, "Failed to parse Wikipedia response", nil, nil)
		return
	}

	// Optional: structure the result into a response map
	data := map[string]interface{}{
		"search_term": result[0],
		"suggestions": result[1],
		"links": result[3],
	}

	response.JSON(w, http.StatusOK, "Success", data, nil)
}
