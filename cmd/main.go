package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	port := ":3000"

	r.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "running..")
	})
	log.Print("server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, r))
}
