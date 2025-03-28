package controllers

import (
	"izzulhaqfs/mnc-tes-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	customers := services.GetCustomers()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Get All Customers Success",
		"data": customers,
	})
}