package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "gitlab.com/rahtz147/cms-coffee-app/app/Database"
	models "gitlab.com/rahtz147/cms-coffee-app/app/Models"
)


func GetWestern(c *gin.Context) {
	var western []models.MenuWestern
	database.DB.Find(&western)
	c.JSON(http.StatusOK, western)
}

func CreateWestern(c *gin.Context)  {
	var westernInput models.MenuWestern
	if err := c.ShouldBindJSON(&westernInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPrice := westernInput.PriceWestern * float64(westernInput.TotalWestern)
	westernInput.TotalPrice = totalPrice

	western := models.MenuWestern{
		NameWestern: westernInput.NameWestern,
		PriceWestern: westernInput.PriceWestern,
		TotalWestern: westernInput.TotalWestern,
		TotalPrice: westernInput.TotalPrice,
	}

	if err := database.DB.Create(&western).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Create Western Menu"})
		return
	}

	database.DB.First(&western)
	c.JSON(http.StatusOK, western)
}

func UpdateWestern(c *gin.Context)  {
	id := c.Param("id")

	var western models.MenuWestern
	if err := database.DB.First(&western, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"The Menu is not updating"})
		return
	}
	if err := c.ShouldBindJSON(&western); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPrice := western.PriceWestern * float64(western.TotalWestern)
	western.TotalPrice = totalPrice

	database.DB.Save(&western)
	c.JSON(http.StatusOK, western)
}

func DeleteWestern(c *gin.Context)  {
	id := c.Param("id")

	var western models.MenuWestern
	if err := database.DB.First(&western, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "The western is not Delete"})
		return
	}



	database.DB.Delete(&western)
	c.JSON(http.StatusOK, gin.H{"message": "western menu Deleted"})
}
