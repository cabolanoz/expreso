package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"expreso-api/config"
	"expreso-api/models"
	"expreso-api/routes"
	"expreso-api/utils"
)

func main() {
	_ = godotenv.Load()

	config.ConnectDB()
	config.DB.AutoMigrate(&models.Admin{}, &models.User{}, &models.Car{})

	port := utils.GetEnv("PORT", "8080")

	r := gin.Default()

	// Register routes
	routes.RegisterUserRoutes(r)

	r.GET("/api/v1/", func(c *gin.Context) {
		c.String(200, "Welcome to Expreso API 😄")
	})

	log.Println("Server running on port", port)
	r.Run(":" + port)
}
