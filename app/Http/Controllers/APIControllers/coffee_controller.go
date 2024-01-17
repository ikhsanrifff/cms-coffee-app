package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "gitlab.com/rahtz147/cms-coffee-app/app/Database"
	models "gitlab.com/rahtz147/cms-coffee-app/app/Models"
)

func GetCoffees(c *gin.Context) {
	var coffees []models.Coffee

	database.DB.Preload("CoffeeType").Find(&coffees)

	c.JSON(http.StatusOK, coffees)
}

func CreateCoffee(c *gin.Context) {
	var coffeeInput models.Coffee
	if err := c.ShouldBindJSON(&coffeeInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPrice := coffeeInput.Price * float64(coffeeInput.TotalCoffee)

	coffeeInput.TotalPrice = totalPrice

	coffee := models.Coffee{
		Name:         coffeeInput.Name,
		Price:        coffeeInput.Price,
		TotalCoffee:  coffeeInput.TotalCoffee,
		TotalPrice:   coffeeInput.TotalPrice,
		CoffeeTypeID: coffeeInput.CoffeeTypeID,
	}

	if err := database.DB.Create(&coffee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create coffee"})
		return
	}

	database.DB.Preload("CoffeeType").First(&coffee, coffee.ID)
	c.JSON(http.StatusOK, coffee)
}

func UpdateCoffee(c *gin.Context) {
	id := c.Param("id")

	var coffee models.Coffee
	if err := database.DB.First(&coffee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
		return
	}

	if err := c.ShouldBindJSON(&coffee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPrice := coffee.Price * float64(coffee.TotalCoffee)
	coffee.TotalPrice = totalPrice

	database.DB.Save(&coffee)
	database.DB.Preload("CoffeeType").First(&coffee)
	c.JSON(http.StatusOK, coffee)
}

func DeleteCoffee(c *gin.Context) {
	id := c.Param("id")

	var coffee models.Coffee
	if err := database.DB.First(&coffee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
		return
	}

	database.DB.Delete(&coffee)

	c.JSON(http.StatusOK, gin.H{"message": "Coffee deleted"})
}

func GetAllCoffees(c *gin.Context) {
	var coffees []models.Coffee
	database.DB.Find(&coffees)
	c.JSON(http.StatusOK, coffees)
}

func GetCoffeeByID(c *gin.Context) {
	id := c.Param("id")

	var coffee models.Coffee
	if err := database.DB.Preload("CoffeeType").First(&coffee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
		return
	}

	c.JSON(http.StatusOK, coffee)
}
