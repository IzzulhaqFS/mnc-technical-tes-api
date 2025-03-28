package models

type Transaction struct {
	Id int `json:"id"`
	SenderId int `json:"sender_id"`
	ReceiverId int `json:"receiver_id"`
	Amount int `json:"amount"`
}