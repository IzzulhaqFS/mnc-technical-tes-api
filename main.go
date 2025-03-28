package main

import (
	"izzulhaqfs/mnc-tes-api/controllers"
	"izzulhaqfs/mnc-tes-api/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	auth := router.Group("/api/auth")

	auth.POST("/login", controllers.Login)

	auth.POST("/logout", controllers.Logout)

	protected := router.Group("/api/protected")
	protected.Use(middlewares.JwtAuthMiddleware())

	protected.POST("/transaction", controllers.CreateTransaction)
	protected.GET("/transaction", controllers.GetTransactions)

	protected.GET("/history", controllers.GetHistories)

	protected.GET("/customer", controllers.GetCustomers)

	router.Run(":8080")
}