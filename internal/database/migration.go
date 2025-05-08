// package database

// import (
// 	"fmt"
// 	"myapps/internal/config"
// )

// func RunMigration() {
// 	migrations := []string{
// 		// Table: users
// 		`CREATE TABLE IF NOT EXISTS users (
// 			id INT AUTO_INCREMENT PRIMARY KEY,
// 			email VARCHAR(255) UNIQUE NOT NULL,
// 			password VARCHAR(255) NOT NULL,
// 			otp VARCHAR(10),
// 			provider VARCHAR(50) NOT NULL DEFAULT 'local',
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			created_by INT,
// 			updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
// 			updated_by INT NULL,
// 			deleted_at TIMESTAMP NULL DEFAULT NULL,
// 			deleted_by INT NULL
// 		);`,

// 		// Table: categories
// 		`CREATE TABLE IF NOT EXISTS categories (
// 			id INT AUTO_INCREMENT PRIMARY KEY,
// 			name VARCHAR(100) NOT NULL,
// 			description TEXT,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			created_by INT,
// 			updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
// 			updated_by INT NULL,
// 			deleted_at TIMESTAMP NULL DEFAULT NULL,
// 			deleted_by INT NULL
// 		);`,

// 		// Table: products
// 		`CREATE TABLE IF NOT EXISTS products (
// 			id INT AUTO_INCREMENT PRIMARY KEY,
// 			name VARCHAR(255) NOT NULL,
// 			description TEXT,
// 			price DECIMAL(12,2) NOT NULL,
// 			stock INT NOT NULL DEFAULT 0,
// 			main_category_id INT,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			created_by INT,
// 			updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
// 			updated_by INT NULL,
// 			deleted_at TIMESTAMP NULL DEFAULT NULL,
// 			deleted_by INT NULL,
// 			FOREIGN KEY (main_category_id) REFERENCES categories(id) ON DELETE SET NULL
// 		);`,

// 		// Table: product_categories (relasi many-to-many)
// 		`CREATE TABLE IF NOT EXISTS product_categories (
// 			product_id INT,
// 			category_id INT,
// 			PRIMARY KEY (product_id, category_id),
// 			FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
// 			FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
// 		);`,

// 		// Table: carts
// 		`CREATE TABLE IF NOT EXISTS carts (
// 			id INT AUTO_INCREMENT PRIMARY KEY,
// 			user_id INT NOT NULL,
// 			product_id INT NOT NULL,
// 			quantity INT NOT NULL DEFAULT 1,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			created_by INT,
// 			updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
// 			updated_by INT NULL,
// 			deleted_at TIMESTAMP NULL DEFAULT NULL,
// 			deleted_by INT NULL,
// 			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
// 			FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
// 		);`,

// 		// Table: transactions
// 		`CREATE TABLE IF NOT EXISTS transactions (
// 			id INT AUTO_INCREMENT PRIMARY KEY,
// 			user_id INT NOT NULL,
// 			total_amount DECIMAL(12,2) NOT NULL,
// 			status ENUM('pending', 'paid', 'shipped', 'completed', 'cancelled') DEFAULT 'pending',
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			created_by INT,
// 			updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
// 			updated_by INT NULL,
// 			deleted_at TIMESTAMP NULL DEFAULT NULL,
// 			deleted_by INT NULL,
// 			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
// 		);`,

// 		// Table: transaction_items
// 		`CREATE TABLE IF NOT EXISTS transaction_items (
// 			id INT AUTO_INCREMENT PRIMARY KEY,
// 			transaction_id INT NOT NULL,
// 			product_id INT NOT NULL,
// 			quantity INT NOT NULL,
// 			price DECIMAL(12,2) NOT NULL,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			created_by INT,
// 			updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
// 			updated_by INT NULL,
// 			deleted_at TIMESTAMP NULL DEFAULT NULL,
// 			deleted_by INT NULL,
// 			FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE CASCADE,
// 			FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
// 		);`,

// 		// Table: reviews
// 		`CREATE TABLE IF NOT EXISTS reviews (
// 			id INT AUTO_INCREMENT PRIMARY KEY,
// 			transaction_item_id INT NOT NULL,
// 			rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5),
// 			comment TEXT,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 			created_by INT,
// 			updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
// 			updated_by INT NULL,
// 			deleted_at TIMESTAMP NULL DEFAULT NULL,
// 			deleted_by INT NULL,
// 			FOREIGN KEY (transaction_item_id) REFERENCES transaction_items(id) ON DELETE CASCADE
// 		);`,
// 	}

// 	// Loop untuk menjalankan semua query
// 	for i, query := range migrations {
// 		_, err := config.DB.Exec(query)
// 		if err != nil {
// 			panic(fmt.Sprintf("❌ Gagal migration ke-%d: %s", i+1, err.Error()))
// 		}
// 	}

// 	fmt.Println("✅ Semua tabel berhasil dimigrasi.")
// }

package database

import (
	"fmt"
	"myapps/internal/config"
)

func RunMigration() {
	migrations := []string{
		// Table: users
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			otp VARCHAR(10),
			provider VARCHAR(50) NOT NULL DEFAULT 'local',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_by INT,
			updated_at TIMESTAMP,
			updated_by INT,
			deleted_at TIMESTAMP,
			deleted_by INT
		);`,

		// Table: categories
		`CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			description TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_by INT,
			updated_at TIMESTAMP,
			updated_by INT,
			deleted_at TIMESTAMP,
			deleted_by INT
		);`,

		// Table: products
		`CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			price NUMERIC(12,2) NOT NULL,
			stock INT NOT NULL DEFAULT 0,
			main_category_id INT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_by INT,
			updated_at TIMESTAMP,
			updated_by INT,
			deleted_at TIMESTAMP,
			deleted_by INT,
			FOREIGN KEY (main_category_id) REFERENCES categories(id) ON DELETE SET NULL
		);`,

		// Table: product_categories
		`CREATE TABLE IF NOT EXISTS product_categories (
			product_id INT,
			category_id INT,
			PRIMARY KEY (product_id, category_id),
			FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
			FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
		);`,

		// Table: carts
		`CREATE TABLE IF NOT EXISTS carts (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			product_id INT NOT NULL,
			quantity INT NOT NULL DEFAULT 1,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_by INT,
			updated_at TIMESTAMP,
			updated_by INT,
			deleted_at TIMESTAMP,
			deleted_by INT,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
		);`,

		// Table: transactions
		`CREATE TABLE IF NOT EXISTS transactions (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			total_amount NUMERIC(12,2) NOT NULL,
			status TEXT DEFAULT 'pending' CHECK (status IN ('pending', 'paid', 'shipped', 'completed', 'cancelled')),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_by INT,
			updated_at TIMESTAMP,
			updated_by INT,
			deleted_at TIMESTAMP,
			deleted_by INT,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,

		// Table: transaction_items
		`CREATE TABLE IF NOT EXISTS transaction_items (
			id SERIAL PRIMARY KEY,
			transaction_id INT NOT NULL,
			product_id INT NOT NULL,
			quantity INT NOT NULL,
			price NUMERIC(12,2) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_by INT,
			updated_at TIMESTAMP,
			updated_by INT,
			deleted_at TIMESTAMP,
			deleted_by INT,
			FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE CASCADE,
			FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
		);`,

		// Table: reviews
		`CREATE TABLE IF NOT EXISTS reviews (
			id SERIAL PRIMARY KEY,
			transaction_item_id INT NOT NULL,
			rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5),
			comment TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_by INT,
			updated_at TIMESTAMP,
			updated_by INT,
			deleted_at TIMESTAMP,
			deleted_by INT,
			FOREIGN KEY (transaction_item_id) REFERENCES transaction_items(id) ON DELETE CASCADE
		);`,
	}

	for i, query := range migrations {
		_, err := config.DB.Exec(query)
		if err != nil {
			panic(fmt.Sprintf("❌ Gagal migration ke-%d: %s", i+1, err.Error()))
		}
	}

	fmt.Println("✅ Semua tabel berhasil dimigrasi.")
}
