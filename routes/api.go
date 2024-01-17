package routes

import (
	"github.com/gin-gonic/gin"
	controllers "gitlab.com/rahtz147/cms-coffee-app/app/Http/Controllers/APIControllers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// CoffeeType routes
	api.GET("/coffee-types", controllers.GetCoffeeTypes)
	api.POST("/coffee-types-add", controllers.CreateCoffeeType)
	api.PUT("/coffee-types-update/:id", controllers.UpdateCoffeeType)
	api.DELETE("/coffee-types-delete/:id", controllers.DeleteCoffeeType)
	api.GET("/coffee-types/:id", controllers.GetCoffeeTypeByID)

	// EmployeeType routes
	api.GET("/employee-types", controllers.GetEmployeeTypes)
	api.POST("/employee-types-add", controllers.CreateEmployeeType)
	api.PUT("/employee-types-update/:id", controllers.UpdateEmployeeType)
	api.DELETE("/employee-types-delete/:id", controllers.DeleteEmployeeType)
	api.GET("/employee-types/:id", controllers.GetEmployeeTypeByID)

	// Coffee routes
	api.GET("/coffees", controllers.GetCoffees)
	api.POST("/coffees-add", controllers.CreateCoffee)
	api.PUT("/coffees-update/:id", controllers.UpdateCoffee)
	api.DELETE("/coffees-delete/:id", controllers.DeleteCoffee)
	api.GET("/coffees/:id", controllers.GetCoffeeByID)

	// Employee routes
	api.GET("/employees", controllers.GetEmployees)
	api.POST("/employees-add", controllers.CreateEmployee)
	api.PUT("/employees-update/:id", controllers.UpdateEmployee)
	api.DELETE("/employees-delete/:id", controllers.DeleteEmployee)
	api.GET("/employees/:id", controllers.GetEmployeeByID)

	//Booking
	api.GET("/bookings", controllers.GetBookings)
	api.POST("/bookings-add", controllers.CreateBooking)
	api.PUT("/bookings-update/:id", controllers.CreateBooking)
	api.DELETE("/bookings-delete/:id", controllers.CreateBooking)
	api.GET("/bookings/:id", controllers.CreateBooking)

	//Table
	api.GET("/tables", controllers.GetTables)
	api.POST("/tables-add", controllers.CreateTable)
	api.PUT("/tables-update", controllers.UpdateTable)
	api.DELETE("/tables-delete", controllers.DeleteTable)
	api.GET("/tables/:id", controllers.GetTableById)

	//History Status
	// api.GET("/histories", controllers.GetBookingHistory)
	api.GET("/:id/histories", controllers.GetBookingHistoriesByID)

	//Absence Employee
	api.GET("/absence-employee", controllers.GetAbsenceEmployee)
	api.POST("/absence-employee-add", controllers.AddAbsenceEmployee)
	api.GET("/absence-employee/:id", controllers.GetAbsenceEmployeeByID)


	//Payment
	api.GET("/payments", controllers.GetPayment)
	api.POST("/payments-add", controllers.CreatePayment)
	api.PUT("/payments-update/:id", controllers.UpdatePayment)
	api.DELETE("/payments-delete/:id", controllers.DeletePayment)
	api.GET("/payments/:id", controllers.GetPaymentById)

	//MENU

	//MenuCake
	api.GET("/cake", controllers.GetCake)
	api.POST("/cake-add", controllers.CreateCake)
	api.PUT("/cake-update/:id", controllers.UpdateCake)
	api.DELETE("/cake-delete/:id", controllers.DeleteCake)

	//MenuSnack
	api.GET("/snack", controllers.GetSnack)
	api.POST("/snack-add", controllers.CreateSnack)
	api.PUT("/snack-update/:id", controllers.UpdateSnack)
	api.DELETE("/snack-delete/:id", controllers.DeleteSnack)

	//MenuWestern
	api.GET("/western", controllers.GetWestern)
	api.POST("/western-add", controllers.CreateWestern)
	api.PUT("/western-update/:id", controllers.UpdateWestern)
	api.DELETE("/western-delete/:id", controllers.DeleteWestern)

	//User
	api.GET("/user", controllers.GetAllUsers)
	api.GET("/user/:id", controllers.GetUserByID)

	//Menu/Foodcategory
	// api.GET("/food-categories", handlers.GetFoodCategories)
	// api.POST("/food-categories", handlers.CreateFoodCategory)
	// api.GET("/foods", handlers.GetFoods)
	// api.POST("/foods", handlers.CreateFood)

}

func SetupAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.POST("/register", controllers.RegisterUser)
	auth.POST("/login", controllers.LoginUser)
}
