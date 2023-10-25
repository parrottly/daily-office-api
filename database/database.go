package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func CreateDBTables(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS year_one_table (
            id SERIAL PRIMARY KEY,
            year VARCHAR(255),
            season VARCHAR(255),
            week VARCHAR(255),
            day VARCHAR(255),
            title VARCHAR(255)
        );
            CREATE TABLE IF NOT EXISTS year_two_table (
            id SERIAL PRIMARY KEY,
            year VARCHAR(255),
            season VARCHAR(255),
            week VARCHAR(255),
            day VARCHAR(255),
            title VARCHAR(255)
        );
            CREATE TABLE IF NOT EXISTS holy_days_table (
            id SERIAL PRIMARY KEY,
            year VARCHAR(255),
            season VARCHAR(255),
            week VARCHAR(255),
            day VARCHAR(255),
            title VARCHAR(255)
        );

            CREATE TABLE IF NOT EXISTS special_occasions_table (
            id SERIAL PRIMARY KEY,
            year VARCHAR(255),
            season VARCHAR(255),
            week VARCHAR(255),
            day VARCHAR(255),
            title VARCHAR(255)
        );
    `)
	if err != nil {
		return err
	}
	return nil
}

func InitDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	var err error
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
