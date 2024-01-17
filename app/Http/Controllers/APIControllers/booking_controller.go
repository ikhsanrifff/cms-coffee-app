package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	database "gitlab.com/rahtz147/cms-coffee-app/app/Database"
	models "gitlab.com/rahtz147/cms-coffee-app/app/Models"
)


func GetBookings(c *gin.Context) {
	var bookings []models.Booking
	database.DB.Preload("User").Preload("Coffee.CoffeeType").Preload("Table").Find(&bookings)
	c.JSON(http.StatusOK, bookings)
}

func CreateBooking(c *gin.Context) {
	var bookingInput models.Booking
	if err := c.ShouldBindJSON(&bookingInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booking := models.Booking{
		NameBooking: bookingInput.NameBooking,
		UserID:      bookingInput.UserID,
		CoffeeID:    bookingInput.CoffeeID,
		TableID:     bookingInput.TableID,
		Schedule:    bookingInput.Schedule,
	}

	if err := database.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	history := models.HistoryBooking{
		BookingID: booking.ID,
		Status:    "Booked",
		Note:      "Booking created",
		UpdatedBy: booking.UserID,
	}
	database.DB.Create(&history)


	database.DB.Preload("User").Preload("Coffee.CoffeeType").Preload("Table").First(&booking, booking.ID)
	c.JSON(http.StatusOK, booking)
}


func UpdateBooking(c *gin.Context)  {
	id := c.Param("id")

	var booking models.Booking
	if err := database.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"The schedule you booked is not available."})
		return
	}
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&booking)
	c.JSON(http.StatusOK, booking)
}

func DeleteBooking(c *gin.Context)  {
	id := c.Param("id")

	var booking models.Booking
	if err := database.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"The schedule you booked is not available."})
		return
	}
	database.DB.Delete(&booking)
	c.JSON(http.StatusOK, gin.H{"message":"Booked deleted"})
}

func GetBookingByID(c *gin.Context)  {
	id := c.Param("id")

	var booking models.Booking
	if err := database.DB.Preload("User").Preload("Coffee").Preload("Table").First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booked not found"})
	}
	c.JSON(http.StatusOK, booking)
}


func GetBookingHistoriesByID(c *gin.Context) {
	bookingID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	var histories []models.HistoryBooking
	database.DB.Where("booking_id = ?", bookingID).Find(&histories)

	database.DB.Preload("Booking.User").Preload("Booking.Coffee.CoffeeType").Preload("Booking.Table").Find(&histories)

	c.JSON(http.StatusOK, histories)
}
