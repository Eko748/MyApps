// routes/routes.go
package routes

import (
	"github.com/rs/cors"
	"net/http"
	"github.com/gorilla/mux"
	"myapps/internal/auth"
	"myapps/internal/controller"
)

type ControllerConfig struct {
	ProductController *controller.ProductController
}

func SetupRoutes(c *ControllerConfig) http.Handler {
	r := mux.NewRouter()

	// Public routes
	route(r, "POST", "/api/register", auth.RegisterHandler)
	route(r, "POST", "/api/login", auth.LoginHandler)

	// API group
	api := r.PathPrefix("/api").Subrouter()

	// Protected routes (activate when needed)
	// api.Use(middleware.AuthMiddleware)

	// Product routes
	route(api, "GET", "/products", c.ProductController.GetProducts)
	route(api, "GET", "/products/{id}", c.ProductController.GetProductByID)
	route(api, "GET", "/products/category/{category_id}", c.ProductController.GetProductsByCategory)
	route(api, "POST", "/products", c.ProductController.CreateProduct)
	route(api, "PUT", "/products", c.ProductController.UpdateProduct)
	route(api, "DELETE", "/products/{id}", c.ProductController.DeleteProduct)
	route(api, "GET", "/products/search", c.ProductController.SearchProducts)
	route(api, "GET", "/products/group/category", c.ProductController.GroupProductsByCategory)
	route(api, "GET", "/products/price-range", c.ProductController.FindProductsByPriceRange)

	// Wrap with CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://ecommerce-web-3zi5dnvew-eko-permanas-projects.vercel.app"}, // Add Vercel domain here
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(r)

	return corsHandler
}

func route(r *mux.Router, method, path string, handlerFunc http.HandlerFunc) {
	r.HandleFunc(path, handlerFunc).Methods(method)
}
