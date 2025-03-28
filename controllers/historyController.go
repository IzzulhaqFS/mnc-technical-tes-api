package controllers

import (
	"izzulhaqfs/mnc-tes-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHistories(c *gin.Context) {
	histories := services.GetHistories()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Get All Histories Success",
		"data": histories,
	})
}