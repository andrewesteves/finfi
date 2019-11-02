package main

import (
	"net/http"

	"github.com/andrewesteves/finfi/controller"
)

func Serve() *http.ServeMux {
	var user controller.UserController
	var client controller.ClientController
	var income controller.IncomeController

	mux := http.NewServeMux()
	mux.HandleFunc("/users", user.Routes())
	mux.HandleFunc("/clients", client.Routes())
	mux.HandleFunc("/incomes", income.Routes())

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	return mux
}
