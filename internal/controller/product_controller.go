package controller

import (
	"encoding/json"
	"myapps/internal/helper"
	"myapps/internal/model"
	"myapps/internal/response"
	"myapps/internal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController struct {
	Service service.ProductService
}

func NewProductController(s service.ProductService) *ProductController {
	return &ProductController{Service: s}
}

// GetProducts retrieves products with pagination
func (c *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
    p := helper.Paginate(r)
    search := r.URL.Query().Get("keyword")

    var (
        products []model.Product
        total    int
        err      error
    )

    if search != "" {
        products, total, err = c.Service.GetPaginatedWithSearch(p.PerPage, p.Offset, search)
        if err != nil {
            response.JSON(w, http.StatusInternalServerError, "Failed to search products", nil, nil)
            return
        }
    } else {
        total, err = c.Service.Count()
        if err != nil {
            response.JSON(w, http.StatusInternalServerError, "Failed to count products", nil, nil)
            return
        }

        products, err = c.Service.GetPaginated(p.PerPage, p.Offset)
        if err != nil {
            response.JSON(w, http.StatusInternalServerError, "Failed to fetch products", nil, nil)
            return
        }
    }

    pagination := response.NewPagination(total, p.PerPage, p.Page)
    response.JSON(w, http.StatusOK, "Success", products, pagination)
}

// GetProductByID retrieves a product by its ID
func (c *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	product, err := c.Service.GetByID(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, "Product not found", nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "Success", product, nil)
}

// GetProductsByCategory retrieves products by category ID
func (c *ProductController) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	categoryID := mux.Vars(r)["category_id"]
	products, err := c.Service.GetByCategory(categoryID)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Failed to fetch products by category", nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "Success", products, nil)
}

// CreateProduct creates a new product
func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		response.JSON(w, http.StatusBadRequest, "Invalid request", nil, nil)
		return
	}
	if err := c.Service.Create(product); err != nil {
		response.JSON(w, http.StatusInternalServerError, "Failed to create product", nil, nil)
		return
	}
	response.JSON(w, http.StatusCreated, "Product created", nil, nil)
}

// UpdateProduct updates an existing product
func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		response.JSON(w, http.StatusBadRequest, "Invalid request", nil, nil)
		return
	}
	if err := c.Service.Update(product); err != nil {
		response.JSON(w, http.StatusInternalServerError, "Failed to update product", nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "Product updated", nil, nil)
}

// DeleteProduct deletes a product by its ID
func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := c.Service.Delete(id); err != nil {
		response.JSON(w, http.StatusInternalServerError, "Failed to delete product", nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "Product deleted", nil, nil)
}

// SearchProducts searches for products based on a query
func (c *ProductController) SearchProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		response.JSON(w, http.StatusBadRequest, "Search query cannot be empty", nil, nil)
		return
	}
	products, err := c.Service.Search(query)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Failed to search products", nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "Success", products, nil)
}

// GroupProductsByCategory groups products by category
func (c *ProductController) GroupProductsByCategory(w http.ResponseWriter, r *http.Request) {
	grouped, err := c.Service.GroupByCategory()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Failed to group products by category", nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "Success", grouped, nil)
}

// FindProductsByPriceRange filters products by price range
func (c *ProductController) FindProductsByPriceRange(w http.ResponseWriter, r *http.Request) {
	minStr := r.URL.Query().Get("min")
	maxStr := r.URL.Query().Get("max")

	min, err1 := strconv.ParseFloat(minStr, 64)
	max, err2 := strconv.ParseFloat(maxStr, 64)
	if err1 != nil || err2 != nil {
		response.JSON(w, http.StatusBadRequest, "Invalid price range", nil, nil)
		return
	}

	products, err := c.Service.FindByPriceRange(min, max)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Failed to filter products by price range", nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "Success", products, nil)
}
