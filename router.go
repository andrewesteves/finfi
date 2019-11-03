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
	mux.HandleFunc("/users", user.Routes())
	mux.HandleFunc("/clients", client.Routes())
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
