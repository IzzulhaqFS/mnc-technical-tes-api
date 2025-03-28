package controllers

import (
	"errors"
	"izzulhaqfs/mnc-tes-api/models"
	"izzulhaqfs/mnc-tes-api/services"
	"izzulhaqfs/mnc-tes-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidateInputAuth struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "Unknown error"
}

func Login(c *gin.Context) {
	var input ValidateInputAuth
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

	var customers models.Customers = services.GetCustomers()
	for _, customer := range customers.Customers {
		if customer.Email == input.Email && customer.Password == input.Password {
			token, err := utils.GenerateToken(customer.Email)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": err.Error(),
				})
				return
			}

			services.AddNewHistory("Login", "Login Success. Email: " + input.Email)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Login Success",
				"data": customer,
				"token": token,
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"message": "Login Failed",
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logout Success",
	})
	services.AddNewHistory("Logout", "Logout Success.")
}