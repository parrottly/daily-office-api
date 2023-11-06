package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"dolapi/database"
	"dolapi/handlers"
)

func main() {
	r := mux.NewRouter()
	port := ":3000"

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Error connecting to database ", err)
	}

	defer db.Close()
	if err := database.CreateDBTables(db); err != nil {
		log.Fatal("Error populating db ", err)
	}

	if err := database.SeedDatabase(db); err != nil {
		log.Fatal("Error populating db ", err)
	}
	r.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "running..")
	})
	r.HandleFunc("/{table}/{season}/{week}/{day}/psalms", func(w http.ResponseWriter, r *http.Request) {
		handlers.PsalmsHandler(w, r, db)
	}).Methods("GET")
	log.Print("server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, r))
}
