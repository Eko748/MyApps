package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    _ "github.com/lib/pq" // Menggunakan driver PostgreSQL
)

var DB *sql.DB

func InitDB() {
    // Membuat connection string untuk PostgreSQL
    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", 
        os.Getenv("DB_USER"), // Username PostgreSQL
        os.Getenv("DB_PASS"), // Password PostgreSQL
        os.Getenv("DB_HOST"), // Host PostgreSQL
        os.Getenv("DB_PORT"), // Port PostgreSQL (default 5432)
        os.Getenv("DB_NAME"), // Nama database PostgreSQL
    )

    var err error
    DB, err = sql.Open("postgres", dsn) // Menggunakan driver PostgreSQL
    if err != nil {
        log.Fatal("Gagal koneksi DB:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("DB tidak merespon:", err)
    }

    fmt.Println("✅ DB connected.")
}


// package config

// import (
//     "database/sql"
//     "fmt"
//     "log"
//     "os"
//     _ "github.com/lib/pq"


//     // _ "github.com/go-sql-driver/mysql"
// )

// var DB *sql.DB

// func InitDB() {
//     dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
//         os.Getenv("DB_USER"),
//         os.Getenv("DB_PASS"),
//         os.Getenv("DB_HOST"),
//         os.Getenv("DB_NAME"),
//     )

//     var err error
//     DB, err = sql.Open("mysql", dsn)
//     if err != nil {
//         log.Fatal("Gagal koneksi DB:", err)
//     }

//     if err = DB.Ping(); err != nil {
//         log.Fatal("DB tidak merespon:", err)
//     }

//     fmt.Println("✅ DB connected.")
// }
