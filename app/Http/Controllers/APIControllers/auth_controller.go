package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "gitlab.com/rahtz147/cms-coffee-app/app/Database"
	models "gitlab.com/rahtz147/cms-coffee-app/app/Models"
)


func RegisterUser(c *gin.Context) {
	var userInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password before saving to database
	if err := userInput.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if err := database.DB.Create(&userInput).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func LoginUser(c *gin.Context) {
	var userInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser models.User
	if err := database.DB.Where("username = ?", userInput.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Check password
	if err := storedUser.CheckPassword(userInput.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Password valid, generate token or set session, etc.
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context)  {
	id := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
