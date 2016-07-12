package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/", index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
