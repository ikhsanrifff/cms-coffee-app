package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "gitlab.com/rahtz147/cms-coffee-app/app/Database"
	models "gitlab.com/rahtz147/cms-coffee-app/app/Models"
)

func GetTables(c *gin.Context) {
	var tables []models.Table
	database.DB.Find(&tables)
	c.JSON(http.StatusOK, tables)
}

func CreateTable(c *gin.Context)  {
	var table models.Table
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&table)
	c.JSON(http.StatusOK, table)
}

func UpdateTable(c *gin.Context)  {
	id := c.Param("id")

	var tables models.Table
	if err := database.DB.First(&tables, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&tables)
	c.JSON(http.StatusOK, tables)
}

func DeleteTable(c *gin.Context)  {
	id := c.Param("id")

	var tables models.Table
	if err := database.DB.First(&tables, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"Table not found"})
	}

	database.DB.Delete(&tables)
	c.JSON(http.StatusOK, gin.H{"message": "Table deleted"})
}

func GetTableById(c *gin.Context)  {
	id := c.Param("id")

	var tables models.Table
	if err := database.DB.First(&tables, id).Error;err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Table not found"})
	}
	c.JSON(http.StatusOK, tables)
}
