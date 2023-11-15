package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"dolapi/internal"
	"dolapi/models"
)

type Lessons struct {
	morning string
	evening string
	gospel  string
}

func LessonsHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	vars := mux.Vars(req)
	tableName := vars["table"]
	season := vars["season"]
	week := vars["week"]
	day := vars["day"]
	weekOfSeason := "Week of " + week + " " + season

	file := internal.GetTable(tableName)
	lessonsData, err := internal.ReadJSONFile(file)
	if err != nil {
		http.Error(resp, "Error reading JSON file", http.StatusInternalServerError)
		return
	}
	if season == "after-pentecost" {
		weekOfSeason = "Proper " + week
	}
	var matchingEntry *models.LiturgicalData
	for _, entry := range lessonsData {
		if strings.EqualFold(entry.Week, weekOfSeason) && strings.EqualFold(entry.Day, day) {
			matchingEntry = &entry
			break
		}
	}

	if matchingEntry == nil {
		http.Error(resp, "Lessons not found", http.StatusNotFound)
		return
	}

	resultJSON, err := json.Marshal(matchingEntry.Lessons)
	if err != nil {
		http.Error(resp, "Error converting result to JSON", http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(resultJSON)
}
