package models

type Customers struct {
	Customers []Customer `json:"customers"`
}

type Customer struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
	Address string `json:"address"`
	Phone string `json:"phone"`
	Balance int `json:"balance"`
}