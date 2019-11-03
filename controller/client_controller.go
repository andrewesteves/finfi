package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/andrewesteves/finfi/model"
	"github.com/gorilla/mux"
)

type ClientController struct{}

func (c ClientController) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clients := model.ClientModel{}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(clients.All())
	}
}

func (c ClientController) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		client := model.ClientModel{
			Name:  fmt.Sprintf("Bill Gates %v", id),
			Email: "bill@microsoft.com",
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(client)
	}
}

func (c ClientController) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var client model.ClientModel
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &client)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(client)
	}
}

func (c ClientController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		client := model.ClientModel{
			Name:  fmt.Sprintf("Bill Gates %v", id),
			Email: "bill@microsoft.com",
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(client)
	}
}

func (c ClientController) Destroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		client := model.ClientModel{
			Name:  fmt.Sprintf("Bill Gates %v", id),
			Email: "bill@microsoft.com",
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(client)
	}
}
