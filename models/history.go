package models

type Histories struct {
	LogActivities []History `json:"histories"`
}

type History struct {
	Id int `json:"id"`
	Action string `json:"action"`
	Detail string `json:"detail"`
}