package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"dolapi/handlers"
)

func main() {
	r := mux.NewRouter()
	port := ":3000"

	r.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "running..")
	})
	r.HandleFunc("/{table}/{season}/{week}/{day}/psalms", func(res http.ResponseWriter, req *http.Request) {
		handlers.PsalmsJSON(res, req)
	}).Methods("GET")
	r.HandleFunc("/{table}/{season}/{week}/{day}/lessons", func(res http.ResponseWriter, req *http.Request) {
		handlers.LessonsHandler(res, req)
	}).Methods("GET")
	r.HandleFunc("/{table}/{season}", func(res http.ResponseWriter, req *http.Request) {
		handlers.SeasonHandler(res, req)
	}).Methods("GET")
	log.Print("server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, r))
}
