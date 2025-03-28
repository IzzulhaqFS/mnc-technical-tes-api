package controllers

import (
	"errors"
	"fmt"
	"izzulhaqfs/mnc-tes-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type validateInputTransaction struct {
	SenderId int `json:"sender_id" binding:"required"`
	ReceiverId int `json:"receiver_id" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}

func CreateTransaction(c *gin.Context) {
	var input validateInputTransaction
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{
					Field:   fe.Field(),
					Message: GetErrorMsg(fe),
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error": out,})
			}
		}
		return
	}

	data, err := services.AddNewTransaction(input.SenderId, input.ReceiverId, input.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	services.AddNewHistory("Transaction", fmt.Sprintf("Sender ID: %d, Receiver ID: %d, Amount: %d", input.SenderId, input.ReceiverId, input.Amount))

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Transaction Success",
		"data": data,
	})
}