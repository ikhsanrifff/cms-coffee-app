package main

import (
	"github.com/gin-gonic/gin"
	database "gitlab.com/rahtz147/cms-coffee-app/app/Database"
)
func main() {
	database.Connect()

    r := gin.Default()

	r.Run(":8080")
}
