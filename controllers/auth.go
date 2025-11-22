package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Secret key para firmar los tokens (en producción usa variable de entorno)
var jwtSecret = []byte("mi_clave_secreta_super_segura")

// Estructura para recibir los datos del login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Handler de login
func Login(c *gin.Context) {
	var req LoginRequest

	// Obtener datos enviados desde Vue
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
		})
		return
	}

	// Credenciales predeterminadas
	if req.Username == "admin" && req.Password == "1234" {
		// Crear el token JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": req.Username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
		})

		// Firmar el token con la clave secreta
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "No se pudo generar el token",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Login exitoso",
			"user":    req.Username,
			"token":   tokenString,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Usuario o contraseña incorrectos",
	})
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Verificamos que el método de firma sea HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return jwtSecret, nil
	})
}
