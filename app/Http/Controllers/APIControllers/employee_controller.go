package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "gitlab.com/rahtz147/cms-coffee-app/app/Database"
	models "gitlab.com/rahtz147/cms-coffee-app/app/Models"
)

func GetEmployees(c *gin.Context) {
	var employees []models.Employee

	database.DB.Preload("EmployeeType").Find(&employees)

	c.JSON(http.StatusOK, employees)
}

func CreateEmployee(c *gin.Context) {
	var employeeInput models.Employee
	if err := c.ShouldBindJSON(&employeeInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee := models.Employee{
		Name: 				employeeInput.Name,
		Age:				employeeInput.Age,
		Address: 			employeeInput.Address,
		Gender: 			employeeInput.Gender,
		EmployeeTypeID: 	employeeInput.EmployeeTypeID,
	}

	if err := database.DB.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"failed to create employee"})
	}

	database.DB.Preload("EmployeeType").First(&employee, employee.ID)
	c.JSON(http.StatusOK, employee)
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee
	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&employee)
	c.JSON(http.StatusOK, employee)
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee
	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	database.DB.Delete(&employee)
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}

func GetAllEmployees(c *gin.Context) {
	var employees []models.Employee
	database.DB.Find(&employees)
	c.JSON(http.StatusOK, employees)
}

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee
	if err := database.DB.Preload("EmployeeType").First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, employee)
}
