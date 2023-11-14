package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"dolapi/internal"
	"dolapi/models"
)

type Psalms struct {
	Morning []string `json:"morning"`
	Evening []string `json:"evening"`
}

func PsalmsHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	vars := mux.Vars(req)
	tableName := vars["table"]
	season := vars["season"]
	week := vars["week"]
	day := vars["day"]
	weekOfSeason := "Week of " + week + " " + season

	file := internal.GetTable(tableName)
	psalmsData, err := internal.ReadJSONFile(file)
	if err != nil {
		http.Error(resp, "Error reading JSON file", http.StatusInternalServerError)
		return
	}

	if season == "after-pentecost" {
		weekOfSeason = "Proper " + week
	}
	var matchingEntry *models.LiturgicalData
	for _, entry := range psalmsData {
		if strings.EqualFold(entry.Week, weekOfSeason) && strings.EqualFold(entry.Day, day) {
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
