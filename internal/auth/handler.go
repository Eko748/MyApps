package auth

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "myapps/internal/config"
    "myapps/internal/model"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var user model.User
    json.NewDecoder(r.Body).Decode(&user)

    _, err := config.DB.Exec("INSERT INTO users (email, password, provider) VALUES (?, ?, ?)",
        user.Email, user.Password, "local")

    if err != nil {
        http.Error(w, "Gagal daftar", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintln(w, "Registrasi berhasil")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var input model.User
    json.NewDecoder(r.Body).Decode(&input)

    row := config.DB.QueryRow("SELECT id, password FROM users WHERE email = ?", input.Email)
    var dbUser model.User
    err := row.Scan(&dbUser.ID, &dbUser.Password)

    if err != nil || dbUser.Password != input.Password {
        http.Error(w, "Login gagal", http.StatusUnauthorized)
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": dbUser.ID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        http.Error(w, "Token error", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
