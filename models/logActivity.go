package models

type LogActivity struct {
	Id int `json:"id"`
	Action string `json:"action"`
	Detail string `json:"detail"`
}