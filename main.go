package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/employee", GetEmployes).Methods("GET")
	r.HandleFunc("/employee/{id}", GetEmployee).Methods("GET")
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employee/{id}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{id}", DeleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitialMigration()
	initializeRouter()
}