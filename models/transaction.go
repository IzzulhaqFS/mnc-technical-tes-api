package models

type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Id int `json:"id"`
	Sender Customer `json:"sender"`
	Receiver Customer `json:"receiver"`
	Amount int `json:"amount"`
}