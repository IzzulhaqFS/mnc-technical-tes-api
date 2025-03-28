package services

import (
	"errors"
	"izzulhaqfs/mnc-tes-api/models"
	"izzulhaqfs/mnc-tes-api/repositories"
)

func GetTransactions() models.Transactions {
	return repositories.GetTransactions()
}

func AddNewTransaction(sender models.Customer, receiverId, amount int) (models.Transaction, error) {
	receiver, err := GetCustomerById(receiverId)
	if err != nil {
		return models.Transaction{}, errors.New("receiver not found")
	}

	if sender.Balance < amount {
		return models.Transaction{}, errors.New("insufficient balance")
	}

	sender.Balance -= amount
	UpdateCustomer(sender)

	receiver.Balance += amount
	UpdateCustomer(receiver)

	return repositories.AddNewTransaction(models.Transaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:     amount,
	}), nil
}