package main

import (
	"net/http"

	"github.com/andrewesteves/finfi/controller"
	"github.com/gorilla/mux"
)

func Serve() *mux.Router {
	var user controller.UserController
	var client controller.ClientController
	var income controller.IncomeController

	mux := mux.NewRouter().StrictSlash(true)
	mux.HandleFunc("/users", user.Index()).Methods("GET")
	mux.HandleFunc("/users/{id}", user.Show()).Methods("GET")
	mux.HandleFunc("/users", user.Store()).Methods("POST")
	mux.HandleFunc("/users/{id}", user.Update()).Methods("PUT")
	mux.HandleFunc("/users/{id}", user.Destroy()).Methods("DELETE")
	mux.HandleFunc("/clients", client.Index()).Methods("GET")
	mux.HandleFunc("/clients", client.Store()).Methods("POST")
	mux.HandleFunc("/clients/{id}", client.Show()).Methods("GET")
	mux.HandleFunc("/clients/{id}", client.Update()).Methods("PUT")
	mux.HandleFunc("/clients/{id}", client.Destroy()).Methods("DELETE")
	mux.HandleFunc("/incomes", income.Index()).Methods("GET")
	mux.HandleFunc("/incomes", income.Store()).Methods("POST")
	mux.HandleFunc("/incomes/{id}", income.Show()).Methods("GET")
	mux.HandleFunc("/incomes/{id}", income.Update()).Methods("PUT")
	mux.HandleFunc("/incomes/{id}", income.Destroy()).Methods("DELETE")

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	return mux
}
