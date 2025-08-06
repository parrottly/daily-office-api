package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"dolapi/internal"
	"dolapi/models"
)

func TodayHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	
	// Get today's liturgical date
	tableName, season, week, day := getTodaysLiturgicalDate()
	
	weekOfSeason := "Week of " + week + " " + season
	if season == "after-pentecost" {
		weekOfSeason = "Proper " + week
	}

	file := internal.GetTable(tableName)
	psalmsData, err := internal.ReadJSONFile(file)
	if err != nil {
		http.Error(resp, "Error reading JSON file", http.StatusInternalServerError)
		return
	}

	var matchingEntry *models.LiturgicalData
	for _, entry := range psalmsData {
		if strings.EqualFold(entry.Week, weekOfSeason) && strings.EqualFold(entry.Day, day) {
			matchingEntry = &entry
			break
		}
	}

	if matchingEntry == nil {
		http.Error(resp, "Today's readings not found", http.StatusNotFound)
		return
	}

	resultJSON, err := json.Marshal(matchingEntry)
	if err != nil {
		http.Error(resp, "Error converting result to JSON", http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(resultJSON)
}

func TodayLessonsHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	
	// Get today's liturgical date
	tableName, season, week, day := getTodaysLiturgicalDate()
	
	weekOfSeason := "Week of " + week + " " + season
	if season == "after-pentecost" {
		weekOfSeason = "Proper " + week
	}

	file := internal.GetTable(tableName)
	psalmsData, err := internal.ReadJSONFile(file)
	if err != nil {
		http.Error(resp, "Error reading JSON file", http.StatusInternalServerError)
		return
	}

	var matchingEntry *models.LiturgicalData
	for _, entry := range psalmsData {
		if strings.EqualFold(entry.Week, weekOfSeason) && strings.EqualFold(entry.Day, day) {
			matchingEntry = &entry
			break
		}
	}

	if matchingEntry == nil {
		http.Error(resp, "Today's readings not found", http.StatusNotFound)
		return
	}

	// Return only the lessons
	resultJSON, err := json.Marshal(matchingEntry.Lessons)
	if err != nil {
		http.Error(resp, "Error converting result to JSON", http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(resultJSON)
}

// Simple liturgical calendar calculation
// This is a basic implementation - in production you'd want more accurate calculations
func getTodaysLiturgicalDate() (string, string, string, string) {
	now := time.Now()
	
	// Get day of week
	dayOfWeek := now.Weekday().String()
	
	// Simple year calculation - alternates between Year One and Year Two
	// Episcopal Church starts Year One in even calendar years, Year Two in odd years
	var tableName string
	if now.Year()%2 == 0 {
		tableName = "year-one"
	} else {
		tableName = "year-two"
	}
	
	// Basic season calculation based on calendar dates
	// This is simplified - actual liturgical calendar is more complex
	month := int(now.Month())
	day := now.Day()
	
	var season, week string
	
	switch {
	case month == 12 && day >= 25:
		season = "christmas"
		week = "1"
	case month == 1 && day <= 6:
		season = "christmas"
		if day <= 31 {
			week = "1"
		}
	case month == 1 && day > 6 || (month == 2):
		season = "epiphany"
		week = "1" // Simplified - should calculate actual week
	case month >= 3 && month <= 4:
		season = "lent"
		week = "1" // Simplified - should calculate actual week based on Easter
	case month == 5:
		season = "easter"
		week = "1" // Simplified - should calculate actual week based on Easter
	case month >= 6 && month <= 11:
		season = "after-pentecost"
		// Simplified calculation for after-pentecost
		weekOfYear := getWeekOfYear(now)
		properWeek := weekOfYear - 20 // Rough approximation
		if properWeek < 1 {
			properWeek = 1
		} else if properWeek > 29 {
			properWeek = 29
		}
		week = intToString(properWeek)
	default:
		season = "advent"
		week = "1"
	}
	
	return tableName, season, week, dayOfWeek
}

func getWeekOfYear(t time.Time) int {
	_, week := t.ISOWeek()
	return week
}

func intToString(i int) string {
	return map[int]string{
		1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "10",
		11: "11", 12: "12", 13: "13", 14: "14", 15: "15", 16: "16", 17: "17", 18: "18", 19: "19", 20: "20",
		21: "21", 22: "22", 23: "23", 24: "24", 25: "25", 26: "26", 27: "27", 28: "28", 29: "29",
	}[i]
}