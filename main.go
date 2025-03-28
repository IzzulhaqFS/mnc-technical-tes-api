package main

import (
	"izzulhaqfs/mnc-tes-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.POST("/login", controllers.Login)

	router.POST("/logout", controllers.Logout)

	router.POST("/transaction", controllers.CreateTransaction)

	router.Run(":8080")
}