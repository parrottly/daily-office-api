package queries

import (
	"database/sql"
	"errors"
	"fmt"
)

func GetPsalms(tableName string, season string, week string, day string, db *sql.DB) ([]string, error) {
	var table string
	weekOfSeason := "Week of " + week + " " + season

	if tableName == "one" {
		table = "year_one_table"
	}
	if tableName == "two" {
		table = "year_two_table"
	}

	if db == nil {
		fmt.Println("Database connection is nil")
		return nil, errors.New("database connection is nil")
	}

	query := "SELECT psalms FROM " + table + " WHERE season = $1 AND week = $2 AND day = $3"
	rows, err := db.Query(query, season, weekOfSeason, day)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}

	defer rows.Close()

	var psalms []string

	for rows.Next() {
		var psalm sql.NullString
		if err := rows.Scan(&psalm); err != nil {
			fmt.Println("Error scanning row:", err)
			break
		}

		if psalm.Valid {
			psalms = append(psalms, psalm.String)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return psalms, nil
}
