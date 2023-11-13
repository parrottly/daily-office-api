package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"dolapi/internal"
	"dolapi/models"
)

type Psalms struct {
	Morning []string `json:"morning"`
	Evening []string `json:"evening"`
}

func PsalmsJSON(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	vars := mux.Vars(req)
	tableName := vars["table"]
	season := vars["season"]
	week := vars["week"]
	day := vars["day"]
	var file string
	weekOfSeason := "Week of " + week + " " + season

	if tableName == "year-one" {
		file = "daily-office/json/readings/dol-year-1.json"
	}
	if tableName == "year-two" {
		file = "daily-office/json/readings/dol-year-2.json"
	}

	psalmsData, err := internal.ReadJSONFile(file)
	if err != nil {
		http.Error(resp, "Error reading JSON file", http.StatusInternalServerError)
		return
	}

	var matchingEntry *models.LiturgicalData
	for _, entry := range psalmsData {
		if entry.Week == weekOfSeason && entry.Day == day {
			matchingEntry = &entry
			break
		}
	}

	if matchingEntry == nil {
		http.Error(resp, "Psalms not found", http.StatusNotFound)
		return
	}

	resultJSON, err := json.Marshal(matchingEntry.Psalms)
	if err != nil {
		http.Error(resp, "Error converting result to JSON", http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(resultJSON)
}
