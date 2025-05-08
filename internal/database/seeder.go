// package database

// import (
// 	"fmt"
// 	"myapps/internal/config"
// )

// func RunSeeder() {
// 	// Seed default user
// 	_, err := config.DB.Exec(`
// 		INSERT IGNORE INTO users (email, password, provider, created_at, created_by)
// 		VALUES ('admin@myapps.com', 'admin123', 'local', CURRENT_TIMESTAMP, 1);
// 	`)
// 	if err != nil {
// 		panic("Seeder gagal: " + err.Error())
// 	}

// 	// Seed categories
// 	_, err = config.DB.Exec(`
// 		INSERT IGNORE INTO categories (id, name, created_at, created_by)
// 		VALUES
// 		(1, 'Elektronik', CURRENT_TIMESTAMP, 1),
// 		(2, 'Pakaian', CURRENT_TIMESTAMP, 1),
// 		(3, 'Makanan', CURRENT_TIMESTAMP, 1);
// 	`)
// 	if err != nil {
// 		panic("Seeder kategori gagal: " + err.Error())
// 	}

// 	// Seed 300 products
// 	productsQuery := `
// 		INSERT IGNORE INTO products (id, name, price, stock, main_category_id, created_at, created_by)
// 		VALUES
// 	`
// 	productValues := ""
// 	for i := 1; i <= 300; i++ { // Ganti dari 30 menjadi 300
// 		name := fmt.Sprintf("'Produk %03d'", i) // Format untuk 3 digit
// 		price := 10000 + (i * 1000)
// 		stock := 10 + (i % 20)
// 		categoryID := 1 + (i % 3) // 1 to 3
// 		productValues += fmt.Sprintf("(%d, %s, %d, %d, %d, CURRENT_TIMESTAMP, 1),\n", i, name, price, stock, categoryID)
// 	}
// 	productsQuery += productValues[:len(productValues)-2] + ";" // remove last comma

// 	_, err = config.DB.Exec(productsQuery)
// 	if err != nil {
// 		panic("Seeder produk gagal: " + err.Error())
// 	}

// 	// Seed product_categories (main_category_id + example many-to-many)
// 	relQuery := `
// 		INSERT IGNORE INTO product_categories (product_id, category_id)
// 		VALUES
// 	`
// 	relValues := ""
// 	for i := 1; i <= 300; i++ { // Ganti dari 30 menjadi 300
// 		c1 := 1 + (i % 3)
// 		c2 := 1 + ((i + 1) % 3)
// 		relValues += fmt.Sprintf("(%d, %d), (%d, %d),\n", i, c1, i, c2)
// 	}
// 	relQuery += relValues[:len(relValues)-2] + ";"

// 	_, err = config.DB.Exec(relQuery)
// 	if err != nil {
// 		panic("Seeder relasi produk-kategori gagal: " + err.Error())
// 	}

// 	fmt.Println("✅ Semua seeder berhasil dijalankan.")
// }

package database

import (
	"fmt"
	"myapps/internal/config"
)

func RunSeeder() {
	// Seed default user
	_, err := config.DB.Exec(`
		INSERT INTO users (id, email, password, provider, created_at, created_by)
		VALUES (1, 'admin@myapps.com', 'admin123', 'local', CURRENT_TIMESTAMP, 1)
		ON CONFLICT (email) DO NOTHING;
	`)
	if err != nil {
		panic("Seeder gagal: " + err.Error())
	}

	// Seed categories
	_, err = config.DB.Exec(`
		INSERT INTO categories (id, name, created_at, created_by)
		VALUES
		(1, 'Elektronik', CURRENT_TIMESTAMP, 1),
		(2, 'Pakaian', CURRENT_TIMESTAMP, 1),
		(3, 'Makanan', CURRENT_TIMESTAMP, 1)
		ON CONFLICT (id) DO NOTHING;
	`)
	if err != nil {
		panic("Seeder kategori gagal: " + err.Error())
	}

	// Seed 300 products
	productsQuery := `
		INSERT INTO products (id, name, price, stock, main_category_id, created_at, created_by)
		VALUES
	`
	productValues := ""
	for i := 1; i <= 300; i++ {
		name := fmt.Sprintf("'Produk %03d'", i)
		price := 10000 + (i * 1000)
		stock := 10 + (i % 20)
		categoryID := 1 + (i % 3)
		productValues += fmt.Sprintf("(%d, %s, %d, %d, %d, CURRENT_TIMESTAMP, 1),\n", i, name, price, stock, categoryID)
	}
	productsQuery += productValues[:len(productValues)-2] + `
	ON CONFLICT (id) DO NOTHING;
	`

	_, err = config.DB.Exec(productsQuery)
	if err != nil {
		panic("Seeder produk gagal: " + err.Error())
	}

	// Seed product_categories
	relQuery := `
		INSERT INTO product_categories (product_id, category_id)
		VALUES
	`
	relValues := ""
	for i := 1; i <= 300; i++ {
		c1 := 1 + (i % 3)
		c2 := 1 + ((i + 1) % 3)
		relValues += fmt.Sprintf("(%d, %d), (%d, %d),\n", i, c1, i, c2)
	}
	relQuery += relValues[:len(relValues)-2] + `
	ON CONFLICT (product_id, category_id) DO NOTHING;
	`

	_, err = config.DB.Exec(relQuery)
	if err != nil {
		panic("Seeder relasi produk-kategori gagal: " + err.Error())
	}

	fmt.Println("✅ Semua seeder berhasil dijalankan.")
}
