package service

import (
	"myapps/internal/model"
	"myapps/internal/repository"
	"myapps/internal/helper"
)

type ProductService interface {
	GetAll() ([]model.Product, error)
	Count() (int, error)
	GetPaginated(limit, offset int) ([]model.Product, error)
	GetPaginatedWithSearch(limit, offset int, search string) ([]model.Product, int, error)
	GetByCategory(string) ([]model.Product, error)
	GetByID(string) (model.Product, error)
	Create(model.Product) error
	Update(model.Product) error
	Delete(string) error
	Search(string) ([]model.Product, error)
	GroupByCategory() (map[string][]model.Product, error)
	FindByPriceRange(float64, float64) ([]model.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{repo: r}
}

func (s *productService) GetAll() ([]model.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) Count() (int, error) {
	return s.repo.Count()
}

func (s *productService) GetPaginated(limit, offset int) ([]model.Product, error) {
    return s.repo.FindPaginated(limit, offset)
}

func (s *productService) GetPaginatedWithSearch(limit, offset int, search string) ([]model.Product, int, error) {
	return s.repo.FindPaginatedWithSearch(limit, offset, search)
}

func (s *productService) GetByCategory(categoryID string) ([]model.Product, error) {
	return s.repo.FindByCategory(categoryID)
}

func (s *productService) GetByID(id string) (model.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) Create(p model.Product) error {
	if err := helper.ValidateProduct(p); err != nil {
		return err
	}
	return s.repo.Create(p)
}

func (s *productService) Update(p model.Product) error {
	if err := helper.ValidateProduct(p); err != nil {
		return err
	}
	return s.repo.Update(p)
}

func (s *productService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *productService) Search(query string) ([]model.Product, error) {
	return s.repo.Search(query)
}

func (s *productService) GroupByCategory() (map[string][]model.Product, error) {
	return s.repo.GroupByCategory()
}

func (s *productService) FindByPriceRange(minPrice, maxPrice float64) ([]model.Product, error) {
	return s.repo.FindByPriceRange(minPrice, maxPrice)
}
