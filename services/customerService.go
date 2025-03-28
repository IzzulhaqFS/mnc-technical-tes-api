package services

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