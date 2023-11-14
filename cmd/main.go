package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"

	"dolapi/handlers"
)

func main() {
	r := mux.NewRouter()
	port := ":80"

	r.Handle("/", templ.Handler(home())).Methods("GET")

	r.HandleFunc("/{table}/{season}/{week}/{day}/psalms", func(res http.ResponseWriter, req *http.Request) {
		handlers.PsalmsHandler(res, req)
	}).Methods("GET")

	r.HandleFunc("/{table}/{season}/{week}/{day}/lessons", func(res http.ResponseWriter, req *http.Request) {
		handlers.LessonsHandler(res, req)
	}).Methods("GET")

	r.HandleFunc("/{table}/{season}/{week}/{day}", func(res http.ResponseWriter, req *http.Request) {
		handlers.DayHandler(res, req)
	}).Methods("GET")

	r.HandleFunc("/{table}/{season}/{week}", func(res http.ResponseWriter, req *http.Request) {
		handlers.WeekHandler(res, req)
	}).Methods("GET")

	r.HandleFunc("/{table}/{season}", func(res http.ResponseWriter, req *http.Request) {
		handlers.SeasonHandler(res, req)
	}).Methods("GET")

	r.HandleFunc("/{table}", func(res http.ResponseWriter, req *http.Request) {
		handlers.YearHandler(res, req)
	}).Methods("GET")

	log.Print("server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, r))
}
