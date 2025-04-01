package controllers

import (
	"expreso-api/config"
	"expreso-api/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetUsers godoc
// @Summary Get all users
// @Description List all registered users with their cars
// @Tags users
// @Produce json
// @Success 200 {object} map[string][]models.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User

	if err := config.DB.Preload("Cars").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	for i := range users {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// CreateUser godoc
// @Summary Create a new user
// @Description Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} map[string]models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
		return
	}
	input.Password = string(hashedPassword)

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	input.Password = ""

	c.JSON(http.StatusCreated, gin.H{"user": input})
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get user details by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]models.User
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.Preload("Cars").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := config.DB.Preload("Cars").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input struct {
		Name         string `json:"name"`
		Email        string `json:"email"`
		PhoneNumber  string `json:"phone_number"`
		ProfilePhoto string `json:"profile_photo"`
		Role         string `json:"role"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	user.PhoneNumber = input.PhoneNumber
	user.ProfilePhoto = input.ProfilePhoto
	user.Role = input.Role

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, gin.H{"user": user})
}
