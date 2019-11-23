package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

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
		param := vars["id"]
		id, err := strconv.Atoi(param)
		if err != nil {
			panic(err.Error())
		}
		client := model.ClientModel{}.Find(id)
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
		cNew, err := client.Insert()
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cNew)
		}
	}
}

func (c ClientController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var client model.ClientModel
		vars := mux.Vars(r)
		param := vars["id"]
		id, _ := strconv.Atoi(param)
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &client)
		w.Header().Add("Content-Type", "application/json")
		cNew, err := client.Update(id)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cNew)
		}
	}
}

func (c ClientController) Destroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var client model.ClientModel
		vars := mux.Vars(r)
		param := vars["id"]
		id, _ := strconv.Atoi(param)
		w.Header().Add("Content-Type", "application/json")
		cNew, err := client.Destroy(id)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cNew)
		}
	}
}
