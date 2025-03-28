package main

import (
	"izzulhaqfs/mnc-tes-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", controllers.Login)

	router.POST("/logout", controllers.Logout)

	router.POST("/transaction", controllers.CreateTransaction)

	router.Run(":8080")
}