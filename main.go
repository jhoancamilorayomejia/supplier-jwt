package main

import (
	"fmt"
	"log"
	"net/http"

	"supplier-jwt/controllers"
	"supplier-jwt/db"

	"github.com/gin-gonic/gin"
)

// Middleware JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
			c.Abort()
			return
		}

		var tokenStr string
		_, err := fmt.Sscanf(authHeader, "Bearer %s", &tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato inválido. Use Bearer <token>"})
			c.Abort()
			return
		}

		token, err := controllers.ValidateToken(tokenStr)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}
	defer database.Close()

	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// ✅ Endpoint público
	r.POST("/api/login", controllers.Login)

	// ✅ Endpoint protegido con JWT
	r.GET("/api/suppliers", AuthMiddleware(), controllers.GetSuppliers)

	log.Println("Servidor corriendo en http://localhost:8080")
	r.Run(":8080")
}
