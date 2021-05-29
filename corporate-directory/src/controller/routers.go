package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func createRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/employee/", getEmployDetails).Methods("GET")
	router.HandleFunc("/employee/{id}", getEmployDetails).Methods("GET")
	router.HandleFunc("/employee/{id}", createEmployee).Methods("POST")
	router.HandleFunc("/employee/{id}", createEmployee).Methods("PUT")
	router.HandleFunc("/employee/{id}", createEmployee).Methods("DELETE")

	router.HandleFunc("/commonmgr/", getEmployDetails).Methods("GET")
	router.HandleFunc("/changemgr/", getEmployDetails).Methods("PUT")
	router.HandleFunc("/moveemployee/", getEmployDetails).Methods("PUT")

}

func createEmployee(w http.ResponseWriter, r *http.Request) {

}

func getEmployDetails(w http.ResponseWriter, r *http.Request) {

}
