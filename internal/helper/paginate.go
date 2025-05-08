package helper

import (
	"net/http"
	"strconv"
)

// Pagination holds pagination values
type Pagination struct {
	Page    int
	PerPage int
	Offset  int
}

// Paginate parses and calculates pagination values from the request
func Paginate(r *http.Request) Pagination {
	pageStr := r.URL.Query().Get("page")
	perPageStr := r.URL.Query().Get("per_page")

	page := 1
	perPage := 10

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if pp, err := strconv.Atoi(perPageStr); err == nil && pp > 0 {
		perPage = pp
	}

	offset := (page - 1) * perPage

	return Pagination{
		Page:    page,
		PerPage: perPage,
		Offset:  offset,
	}
}
