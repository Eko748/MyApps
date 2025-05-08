package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASS"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_NAME"),
    )

    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Gagal koneksi DB:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("DB tidak merespon:", err)
    }

    fmt.Println("✅ DB connected.")
}
