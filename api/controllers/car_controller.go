package controllers

import (
	"expreso-api/config"
	"expreso-api/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// AddUserCar godoc
// @Summary Create a new car for a user
// @Description Register a new car for a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param car body models.Car true "Car data"
// @Success 201 {object} map[string]models.Car
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id}/cars [post]
func AddUserCar(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserId = user.ID

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create car"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"car": input})
}

// UpdateUserCar godoc
// @Summary Update a car for a user
// @Description Update a car for a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param car_id path int true "Car ID"
// @Param car body models.Car true "Car data"
// @Success 200 {object} map[string]models.Car
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id}/cars/{car_id} [put]
func UpdateUserCar(c *gin.Context) {
	userID := c.Param("id")
	carID := c.Param("car_id")

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var car models.Car
	if err := config.DB.Where("id = ? AND user_id = ?", carID, userID).First(&car).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found for this user"})
		return
	}

	var input struct {
		Make    string `json:"make"`
		Pattern string `json:"pattern"`
		Year    int    `json:"year"`
		Plate   string `json:"plate"`
		Color   string `json:"color"`
		Seats   int    `json:"seats"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car.Make = input.Make
	car.Pattern = input.Pattern
	car.Year = input.Year
	car.Plate = input.Plate
	car.Color = input.Color
	car.Seats = input.Seats

	if err := config.DB.Save(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"car": car})
}

// DeleteUserCar godoc
// @Summary Delete a car for a user
// @Description Delete a car for a user
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Param car_id path int true "Car ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id}/cars/{car_id} [delete]
func DeleteUserCar(c *gin.Context) {
	userID := c.Param("id")
	carID := c.Param("car_id")

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var car models.Car
	if err := config.DB.Where("id = ? AND user_id = ?", carID, userID).First(&car).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found for this user"})
		return
	}

	if err := config.DB.Delete(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete car"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}

