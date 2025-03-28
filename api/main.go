package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"expreso-api/config"
	"expreso-api/models"
	"expreso-api/utils"
)

func main() {
	_ = godotenv.Load()

	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	port := utils.GetEnv("PORT", "8080")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Expreso API 👋",
		})
	})

	log.Println("Server running on port", port)
	r.Run(":" + port)
}
