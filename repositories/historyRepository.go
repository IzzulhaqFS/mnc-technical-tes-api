package repositories

import (
	"encoding/json"
	"izzulhaqfs/mnc-tes-api/models"
	"log"
	"os"
)

func GetHistories() models.Histories {
	var histories models.Histories

	jsonFile, err := os.ReadFile("data/histories.json")
	if err != nil {
		log.Fatalf("error when opening file: %v", err)
	}

	json.Unmarshal(jsonFile, &histories)

	return histories
}

func AddNewHistory(history models.History) models.History {
	histories := GetHistories()

	history.Id = len(histories.LogActivities) + 1

	histories.LogActivities = append(histories.LogActivities, history)
	data, _ := json.Marshal(histories)
	os.WriteFile("data/histories.json", data, 0644)
	return history
}