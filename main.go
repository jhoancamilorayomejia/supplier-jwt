package main

import (
	"log"
	"supplier-jwt/controllers"
	"supplier-jwt/db"

	"github.com/gin-gonic/gin"
)

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
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/api/login", controllers.Login)

	//para obtener la informacion de los proveedores
	r.GET("/api/suppliers", controllers.GetSuppliers)

	r.Run(":8080")
}
