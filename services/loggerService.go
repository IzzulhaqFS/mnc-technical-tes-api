package services

import (
	"encoding/json"
	"izzulhaqfs/mnc-tes-api/models"
	"log"
	"os"
)

func GetLogs() []models.LogActivity {
	var logEntries []models.LogActivity

	jsonFile, err := os.ReadFile("data/history.json")
	if err != nil {
		log.Fatalf("error when opening file: %v", err)
	}

	json.Unmarshal(jsonFile, &logEntries)

	return logEntries
}

func LogActivity(action string, detail string) {
	logEntries := GetLogs()
	logEntries = append(logEntries, models.LogActivity{Action: action, Detail: detail})
	data, _ := json.Marshal(logEntries)
	os.WriteFile("data/history.json", data, 0644)
}