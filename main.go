package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type resError struct {
	Type        int    `json:"type"`
	Description string `json:"description"`
	Level       int    `json:"level"`
}

type meta struct {
	Records int    `json:"records"`
	Type    string `json:"type"`
	Time    string `json:"time"`
}

type ocResponse struct {
	Body  interface{} `json:"body"`
	Error resError    `json:"error"`
	Meta  meta        `json:"meta"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/", indexHandler)
	router.HandleFunc("/v1/user", userHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
