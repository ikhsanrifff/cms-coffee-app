package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "gitlab.com/rahtz147/cms-coffee-app/app/Database"
	models "gitlab.com/rahtz147/cms-coffee-app/app/Models"
)


func GetEmployeeTypes(c *gin.Context) {
	var employeeTypes []models.EmployeeType
	database.DB.Find(&employeeTypes)
	c.JSON(http.StatusOK, employeeTypes)
}

func CreateEmployeeType(c *gin.Context) {
	var employeeType models.EmployeeType
	if err := c.ShouldBindJSON(&employeeType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&employeeType)
	c.JSON(http.StatusOK, employeeType)
}

func UpdateEmployeeType(c *gin.Context) {
	id := c.Param("id")

	var employeeType models.EmployeeType
	if err := database.DB.First(&employeeType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "EmployeeType not found"})
		return
	}

	if err := c.ShouldBindJSON(&employeeType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&employeeType)
	c.JSON(http.StatusOK, employeeType)
}

func DeleteEmployeeType(c *gin.Context) {
	id := c.Param("id")

	var employeeType models.EmployeeType
	if err := database.DB.First(&employeeType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "EmployeeType not found"})
		return
	}

	database.DB.Delete(&employeeType)
	c.JSON(http.StatusOK, gin.H{"message": "EmployeeType deleted"})
}

func GetAllEmployeeTypes(c *gin.Context) {
	var employeeTypes []models.EmployeeType
	database.DB.Find(&employeeTypes)
	c.JSON(http.StatusOK, employeeTypes)
}

func GetEmployeeTypeByID(c *gin.Context) {
	id := c.Param("id")

	var employeeType models.EmployeeType
	if err := database.DB.First(&employeeType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "EmployeeType not found"})
		return
	}

	c.JSON(http.StatusOK, employeeType)
}
