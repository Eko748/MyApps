package repository

import (
	"database/sql"
	"log"
	"myapps/internal/config"
	"myapps/internal/helper"
	"myapps/internal/model"
	"strconv"
)

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	Count() (int, error)
	FindPaginated(limit, offset int) ([]model.Product, error)
	FindByID(id string) (model.Product, error)
	FindByCategory(categoryID string) ([]model.Product, error)
	FindByName(name string) ([]model.Product, error)
	FindByPriceRange(minPrice, maxPrice float64) ([]model.Product, error)
	Create(product model.Product) error
	Update(product model.Product) error
	Delete(id string) error
	Search(query string) ([]model.Product, error)
	GroupByCategory() (map[string][]model.Product, error)
}

type productRepo struct {
	db *sql.DB
}

func NewProductRepository() ProductRepository {
	return &productRepo{db: config.DB}
}

// FindAll retrieves all products
func (r *productRepo) FindAll() ([]model.Product, error) {
	query := `SELECT id, name, description, price, stock, main_category_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM products`
	rows, err := helper.ExecuteQuery(r.db, query)
	if err != nil {
		log.Println("Query error:", err) // Tambahkan ini untuk debug
		return nil, err
	}

	defer rows.Close()

	return helper.ProductRows(rows)
}

// Count returns the total number of products (excluding deleted ones)
func (r *productRepo) Count() (int, error) {
	query := `SELECT COUNT(*) FROM products WHERE deleted_at IS NULL`
	row := r.db.QueryRow(query)

	var total int
	err := row.Scan(&total)
	return total, err
}

// FindPaginated retrieves products using limit and offset
func (r *productRepo) FindPaginated(limit, offset int) ([]model.Product, error) {
    query := `SELECT id, name, description, price, stock, main_category_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
				FROM products
				WHERE deleted_at IS NULL
				ORDER BY created_at DESC
				LIMIT ? OFFSET ?`

    rows, err := r.db.Query(query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return helper.ProductRows(rows)
}


// FindByID retrieves a product by ID
func (r *productRepo) FindByID(id string) (model.Product, error) {
	query := `SELECT id, name, description, price, stock, main_category_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM products WHERE id = ?`
	rows, err := helper.ExecuteQuery(r.db, query, id)
	if err != nil {
		return model.Product{}, err
	}
	defer rows.Close()

	products, err := helper.ProductRows(rows)
	if err != nil {
		return model.Product{}, err
	}

	if len(products) > 0 {
		return products[0], nil
	}
	return model.Product{}, nil
}

// FindByCategory retrieves products by category ID
func (r *productRepo) FindByCategory(categoryID string) ([]model.Product, error) {
	query := `SELECT id, name, description, price, stock, main_category_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM products WHERE main_category_id = ?`
	rows, err := helper.ExecuteQuery(r.db, query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return helper.ProductRows(rows)
}

// FindByName retrieves products by name (partial match)
func (r *productRepo) FindByName(name string) ([]model.Product, error) {
	query := `SELECT id, name, description, price, stock, main_category_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM products WHERE name LIKE ?`
	rows, err := helper.ExecuteQuery(r.db, query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return helper.ProductRows(rows)
}

// FindByPriceRange retrieves products by price range
func (r *productRepo) FindByPriceRange(minPrice, maxPrice float64) ([]model.Product, error) {
	query := `SELECT id, name, description, price, stock, main_category_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM products WHERE price BETWEEN ? AND ?`
	rows, err := helper.ExecuteQuery(r.db, query, minPrice, maxPrice)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return helper.ProductRows(rows)
}

// Create inserts a new product
func (r *productRepo) Create(p model.Product) error {
	query := `INSERT INTO products (name, description, price, stock, main_category_id, created_by) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, p.Name, p.Description, p.Price, p.Stock, p.MainCategoryID, p.CreatedBy)
	return err
}

// Update modifies an existing product
func (r *productRepo) Update(p model.Product) error {
	query := `UPDATE products SET name = ?, description = ?, price = ?, stock = ?, main_category_id = ?, updated_at = ?, updated_by = ? WHERE id = ?`
	_, err := r.db.Exec(query, p.Name, p.Description, p.Price, p.Stock, p.MainCategoryID, p.UpdatedAt, p.UpdatedBy, p.ID)
	return err
}

// Delete soft deletes a product by setting deleted_at and deleted_by
func (r *productRepo) Delete(id string) error {
	query := `UPDATE products SET deleted_at = NOW(), deleted_by = ? WHERE id = ?`
	_, err := r.db.Exec(query, 1, id) // assuming `1` is the ID of the user performing the delete
	return err
}

// Search searches for products based on a search query
func (r *productRepo) Search(query string) ([]model.Product, error) {
	queryStr := `%` + query + `%`
	return r.FindByName(queryStr)
}

// GroupByCategory groups products by their main category
func (r *productRepo) GroupByCategory() (map[string][]model.Product, error) {
	query := `SELECT id, name, description, price, stock, main_category_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM products`
	rows, err := helper.ExecuteQuery(r.db, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	grouped := make(map[string][]model.Product)

	products, err := helper.ProductRows(rows)
	if err != nil {
		return nil, err
	}

	// Group products by category
	for _, p := range products {
		categoryID := strconv.Itoa(p.MainCategoryID)
		grouped[categoryID] = append(grouped[categoryID], p)
	}

	return grouped, nil
}
