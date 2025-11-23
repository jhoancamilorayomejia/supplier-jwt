package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"supplier-jwt/db" // tu conexión a la DB

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Clave secreta para JWT
var jwtSecret = []byte("mi_clave_secreta_super_segura")

// Estructura para recibir datos del login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Estructura de usuario desde la DB
type User struct {
	Username string
	Password string
	Rol      string
}

// Handler de login
func Login(c *gin.Context) {
	var req LoginRequest

	// Leer JSON enviado desde Vue
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Buscar usuario en la DB
	var user User
	err := db.DB.QueryRow("SELECT username, password, rol FROM \"user\" WHERE username=$1", req.Username).
		Scan(&user.Username, &user.Password, &user.Rol)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Verificar contraseña directamente
	if req.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	// Crear token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"rol":      user.Rol,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"user":    user.Username,
		"rol":     user.Rol,
		"token":   tokenString,
	})
}

// Función para validar JWT
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
}
