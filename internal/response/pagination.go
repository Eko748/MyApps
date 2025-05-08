package response

import "math"

type Pagination struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}

func NewPagination(total, perPage, currentPage int) Pagination {
	totalPages := int(math.Ceil(float64(total) / float64(perPage)))
	return Pagination{
		Total:       total,
		PerPage:     perPage,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
	}
}
