package routes

import (
	controller "expreso-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	// User routes
	userRoutes := r.Group("/api/v1/users")
	{
		userRoutes.GET("/", controller.GetUsers)
		userRoutes.POST("/", controller.CreateUser)
		userRoutes.GET("/:id", controller.GetUserByID)
		userRoutes.PUT("/:id", controller.UpdateUser)
		userRoutes.POST("/:id/cars", controller.AddUserCar)
		userRoutes.PUT("/:id/cars/:car_id", controller.UpdateUserCar)
		userRoutes.DELETE("/:id/cars/:car_id", controller.DeleteUserCar)
	}
}
