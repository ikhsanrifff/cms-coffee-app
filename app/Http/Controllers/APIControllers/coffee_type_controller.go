package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "gitlab.com/rahtz147/cms-coffee-app/app/Database"
	models "gitlab.com/rahtz147/cms-coffee-app/app/Models"
)



func GetCoffeeTypes(c *gin.Context) {
	var coffeeTypes []models.CoffeeType
	database.DB.Find(&coffeeTypes)
	c.JSON(http.StatusOK, coffeeTypes)
}

func CreateCoffeeType(c *gin.Context) {
	var coffeeType models.CoffeeType
	if err := c.ShouldBindJSON(&coffeeType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&coffeeType)
	c.JSON(http.StatusOK, coffeeType)
}

func UpdateCoffeeType(c *gin.Context) {
	id := c.Param("id")
	var coffeeType models.CoffeeType
	if err := database.DB.First(&coffeeType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CoffeeType not found"})
		return
	}

	if err := c.ShouldBindJSON(&coffeeType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&coffeeType)
	c.JSON(http.StatusOK, coffeeType)
}

func DeleteCoffeeType(c *gin.Context) {
	id := c.Param("id")

	var coffeeType models.CoffeeType
	if err := database.DB.First(&coffeeType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CoffeeType not found"})
		return
	}

	database.DB.Delete(&coffeeType)
	c.JSON(http.StatusOK, gin.H{"message": "CoffeeType deleted"})
}

func GetAllCoffeeTypes(c *gin.Context) {
	var coffeeTypes []models.CoffeeType
	database.DB.Find(&coffeeTypes)
	c.JSON(http.StatusOK, coffeeTypes)
}

func GetCoffeeTypeByID(c *gin.Context) {
	id := c.Param("id")

	var coffeeType models.CoffeeType
	if err := database.DB.First(&coffeeType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CoffeeType not found"})
		return
	}

	c.JSON(http.StatusOK, coffeeType)
}
