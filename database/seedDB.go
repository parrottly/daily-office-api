package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"dolapi/models"
)

func SeedDatabase(db *sql.DB) error {
	jsonFiles := [4]string{
		"daily-office/json/readings/dol-year-1.json",
		"daily-office/json/readings/dol-year-2.json",
		"daily-office/json/readings/dol-holy-days.json",
		"daily-office/json/readings/dol-special-occasions.json",
	}
	tableNames := [4]string{
		"year_one_table",
		"year_two_table",
		"holy_days_table",
		"special_occasions_table",
	}

	for i, jsonFile := range jsonFiles {
		data, err := os.ReadFile(jsonFile)
		if err != nil {
			log.Fatal(err)
		}
		var jsonData []models.LiturgicalData
		if err := json.Unmarshal(data, &jsonData); err != nil {
			log.Fatal(err)
		}
		tablename := tableNames[i]
		for _, dol := range jsonData {
			_, err := db.Exec("INSERT INTO "+tablename+" (year, season, week, day, title) VALUES ($1, $2, $3, $4, $5)",
				dol.Year, dol.Season, dol.Week, dol.Day, dol.Title)
			if err != nil {
				fmt.Printf("Failed to insert data into %s: %v\n", jsonFile)
			}
		}
		fmt.Printf("Data import completed for %s.\n", jsonFile)
	}
	return nil
}
