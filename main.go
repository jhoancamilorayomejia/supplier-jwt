package main

import (
	"supplier-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Habilitar CORS simple
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

	r.Run(":8080")
}
