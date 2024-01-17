package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "gitlab.com/rahtz147/cms-coffee-app/app/Database"
	models "gitlab.com/rahtz147/cms-coffee-app/app/Models"
)


func GetSnack(c *gin.Context) {
	var snack []models.MenuSnack
	database.DB.Find(&snack)
	c.JSON(http.StatusOK, snack)
}

func CreateSnack(c *gin.Context)  {
	var snackInput models.MenuSnack
	if err := c.ShouldBindJSON(&snackInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPrice := snackInput.PriceSnack * float64(snackInput.TotalSnack)
	snackInput.TotalPrice = totalPrice

	snack := models.MenuSnack{
		NameSnack: snackInput.NameSnack,
		PriceSnack: snackInput.PriceSnack,
		TotalSnack: snackInput.TotalSnack,
		TotalPrice: snackInput.TotalPrice,
	}

	if err := database.DB.Create(&snack).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Create Snack Menu"})
		return
	}

	database.DB.First(&snack, snack.ID)
	c.JSON(http.StatusOK, snack)
}

func UpdateSnack(c *gin.Context)  {
	id := c.Param("id")

	var snack models.MenuSnack
	if err := database.DB.First(&snack, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"The Menu is not updating"})
		return
	}
	if err := c.ShouldBindJSON(&snack); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPrice := snack.PriceSnack * float64(snack.TotalSnack)
	snack.TotalPrice = totalPrice

	database.DB.Save(&snack)
	c.JSON(http.StatusOK, snack)
}

func DeleteSnack(c *gin.Context)  {
	id := c.Param("id")

	var snack models.MenuSnack
	if err := database.DB.First(&snack, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "The snack is not Delete"})
		return
	}
	database.DB.Delete(&snack)
	c.JSON(http.StatusOK, gin.H{"message": "Cake Deleted"})
}
