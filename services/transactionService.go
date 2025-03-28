package services

import (
	"izzulhaqfs/mnc-tes-api/models"
	"log"
	"os"
	"encoding/json"
)

func GetTransactions() []models.Transaction {
	var transactions []models.Transaction

	jsonFile, err := os.ReadFile("data/transactions.json")
	if err != nil {
		log.Fatalf("error when opening file: %v", err)
	}

	json.Unmarshal(jsonFile, &transactions)

	return transactions
}

func AddNewTransaction(transaction models.Transaction) {
	transactions := GetTransactions()
	transactions = append(transactions, transaction)
	data, _ := json.Marshal(transactions)
	os.WriteFile("data/transactions.json", data, 0644)
}