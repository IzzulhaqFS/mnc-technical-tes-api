package services

import (
	"encoding/json"
	"izzulhaqfs/mnc-tes-api/models"
	"log"
	"os"
)

func GetLogs() models.LogActivities {
	var logEntries models.LogActivities

	jsonFile, err := os.ReadFile("data/history.json")
	if err != nil {
		log.Fatalf("error when opening file: %v", err)
	}

	json.Unmarshal(jsonFile, &logEntries)

	return logEntries
}

func LogActivity(action string, detail string) {
	logEntries := GetLogs()
	logEntries.LogActivities = append(logEntries.LogActivities, models.LogActivity{Id: len(logEntries.LogActivities) + 1, Action: action, Detail: detail})
	data, _ := json.Marshal(logEntries)
	os.WriteFile("data/history.json", data, 0644)
}