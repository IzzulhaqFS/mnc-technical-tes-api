package services

import (
	"izzulhaqfs/mnc-tes-api/models"
	"log"
	"os"
	"encoding/json"
)

func GetTransactions() models.Transactions {
	var transactions models.Transactions

	jsonFile, err := os.ReadFile("data/transactions.json")
	if err != nil {
		log.Fatalf("error when opening file: %v", err)
	}

	json.Unmarshal(jsonFile, &transactions)

	return transactions
}

func AddNewTransaction(transaction models.Transaction) models.Transaction {
	transactions := GetTransactions()

	transaction.Id = len(transactions.Transactions) + 1

	transactions.Transactions = append(transactions.Transactions, transaction)
	data, _ := json.Marshal(transactions)
	os.WriteFile("data/transactions.json", data, 0644)
	return transaction
}