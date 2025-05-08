package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
    "myapps/internal/config"
    "myapps/internal/database"
    "myapps/internal/controller"
    "myapps/internal/repository"
    "myapps/internal/routes"
    "myapps/internal/service"
)

func main() {
    // Load environment variables
    godotenv.Load()

    // Initialize database connection
    config.InitDB()

    // Run migrations and seeding
    database.RunMigration()
    database.RunSeeder()

    // Initialize repositories
    productRepo := repository.NewProductRepository()
    // transactionRepo := repository.NewTransactionRepository(config.DB)
    // cartRepo := repository.NewCartRepository(config.DB)
    // userRepo := repository.NewUserRepository(config.DB)

    // Initialize services
    productService := service.NewProductService(productRepo)
    // transactionService := service.NewTransactionService(transactionRepo)
    // cartService := service.NewCartService(cartRepo)
    // userService := service.NewUserService(userRepo)

    // Initialize controllers
    productController := controller.NewProductController(productService)
    // transactionController := controller.NewTransactionController(transactionService)
    // cartController := controller.NewCartController(cartService)
    // userController := controller.NewUserController(userService)

    // Create ControllerConfig instance
    controllers := &routes.ControllerConfig{
        ProductController:     productController,
        // TransactionController: transactionController,
        // CartController:        cartController,
        // UserController:        userController,
    }

    // Setup routes
    r := routes.SetupRoutes(controllers)

    // Get the server port from environment variable or default to 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Start the server
    fmt.Println("Server running on port:", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
