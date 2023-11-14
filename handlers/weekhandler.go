package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"dolapi/internal"
	"dolapi/models"
)

func WeekHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	vars := mux.Vars(req)
	tableName := vars["table"]
	season := vars["season"]
	week := vars["week"]
	weekOfSeason := "Week of " + week + " " + season

	file := internal.GetTable(tableName)
	weekData, err := internal.ReadJSONFile(file)
	if err != nil {
		http.Error(resp, "Error reading JSON file", http.StatusInternalServerError)
		return
	}

	matchingEntries := []models.LiturgicalData{}
	for _, entry := range weekData {
		if strings.EqualFold(entry.Week, weekOfSeason) {
			matchingEntries = append(matchingEntries, entry)
		}
	}

	if matchingEntries == nil {
		http.Error(resp, "Week not found", http.StatusNotFound)
		return
	}

	resultJSON, err := json.Marshal(matchingEntries)
	if err != nil {
		http.Error(resp, "Error converting result to JSON", http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(resultJSON)
}
