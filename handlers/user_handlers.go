package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"go-books/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key")

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func UserLogin(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Cek user dari DB
    var dbPassword string
    err := config.DB.QueryRow("SELECT password FROM users WHERE username=$1", req.Username).Scan(&dbPassword)
    if err == sql.ErrNoRows {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // ⚠️ Sementara plain text (nanti sebaiknya pakai bcrypt)
    if req.Password != dbPassword {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        return
    }

    // Buat token JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": req.Username,
        "exp":      time.Now().Add(time.Hour).Unix(),
    })

    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
