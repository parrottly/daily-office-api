package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"dolapi/internal"
	"dolapi/models"
)

func SeasonHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	vars := mux.Vars(req)
	tableName := vars["table"]
	season := vars["season"]

	file := internal.GetTable(tableName)
	seasonData, err := internal.ReadJSONFile(file)
	if err != nil {
		http.Error(resp, "Error reading JSON file", http.StatusInternalServerError)
		return
	}

	if season == "after-pentecost" {
		season = "The Season After Pentecost"
	}
	matchingEntries := []models.LiturgicalData{}
	for _, entry := range seasonData {
		if strings.EqualFold(entry.Season, season) {
			matchingEntries = append(matchingEntries, entry)
		}
	}

	if matchingEntries == nil {
		http.Error(resp, "Season not found", http.StatusNotFound)
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
