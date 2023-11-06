package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"dolapi/queries"
)

func PsalmsHandler(resp http.ResponseWriter, req *http.Request, db *sql.DB) {
	resp.Header().Set("Content-type", "application/json")

	if db == nil {
		fmt.Println("db is nil")
	}

	vars := mux.Vars(req)
	tablename := vars["table"]
	season := vars["season"]
	week := vars["week"]
	day := vars["day"]

	psalms, err := queries.GetPsalms(tablename, season, week, day, db)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		resp.Write([]byte(`{"error": "Error getting psalms"}`))
		return
	}
	jsonResp, err := json.Marshal(psalms)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(jsonResp)
}
