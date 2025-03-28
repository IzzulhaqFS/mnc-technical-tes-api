package repositories

import (
	"encoding/json"
	"izzulhaqfs/mnc-tes-api/models"
	"log"
	"os"
)

func GetCustomers() models.Customers {
	var customers models.Customers

	jsonFile, err := os.ReadFile("data/customers.json")
	if err != nil {
		log.Fatalf("error when opening file: %v", err)
	}

	json.Unmarshal(jsonFile, &customers)

	return customers
}

func GetCustomerById(id int) models.Customer {
	customers := GetCustomers()

	for _, customer := range customers.Customers {
		if customer.Id == id {
			return customer
		}
	}

	return models.Customer{}
}

func UpdateCustomer(customer models.Customer) models.Customer {
	customers := GetCustomers()

	for i, c := range customers.Customers {
		if c.Id == customer.Id {
			customers.Customers[i] = customer
			break
		}
	}

	data, _ := json.Marshal(customers)
	os.WriteFile("data/customers.json", data, 0644)

	return customer
}