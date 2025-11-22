package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		c.JSON(http.StatusOK, gin.H{
			"message": "Login exitoso",
			"user":    req.Username,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Usuario o contraseña incorrectos",
	})
}
