package models

type LogActivities struct {
	LogActivities []LogActivity `json:"history"`
}

type LogActivity struct {
	Id int `json:"id"`
	Action string `json:"action"`
	Detail string `json:"detail"`
}