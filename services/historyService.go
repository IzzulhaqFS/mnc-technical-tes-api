package services

import (
	"izzulhaqfs/mnc-tes-api/models"
	"izzulhaqfs/mnc-tes-api/repositories"
)

func GetHistories() models.Histories {
	return repositories.GetHistories()
}

func AddNewHistory(action string, detail string) models.History {
	return repositories.AddNewHistory(models.History{Action: action, Detail: detail})
}