package routes

import (
	"expreso-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	// User routes
	userRoutes := r.Group("/api/v1/users")
	{
		userRoutes.POST("/", controllers.Create)
	}
}
