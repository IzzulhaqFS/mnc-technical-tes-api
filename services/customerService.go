package services

import (
	"errors"
	"izzulhaqfs/mnc-tes-api/models"
	"izzulhaqfs/mnc-tes-api/repositories"
)

func GetCustomers() models.Customers {
	return repositories.GetCustomers()
}

func GetCustomerById(id int) (models.Customer, error) {
	customer := repositories.GetCustomerById(id)
	if (customer == models.Customer{}) {
		return models.Customer{}, errors.New("customer not found")
	}
	return customer, nil
}

func GetCustomerByEmail(email string) (models.Customer, error) {
	customer := repositories.GetCustomerByEmail(email)
	if (customer == models.Customer{}) {
		return models.Customer{}, errors.New("customer not found")
	}
	return customer, nil
}

func UpdateCustomer(customer models.Customer) models.Customer {
	return repositories.UpdateCustomer(customer)
}