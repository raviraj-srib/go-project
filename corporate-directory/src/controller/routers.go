package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func createRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/employee", createEmployee).Methods("POST")
	router.HandleFunc("/employee/{id}", getEmployDetails).Methods("GET")

}

func createEmployee(w http.ResponseWriter, r *http.Request) {

}

func getEmployDetails(w http.ResponseWriter, r *http.Request) {

}
