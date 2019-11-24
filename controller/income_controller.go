package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/andrewesteves/finfi/model"
	"github.com/gorilla/mux"
)

type IncomeController struct{}

func (i IncomeController) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		incomes := model.IncomeModel{}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(incomes.All())
	}
}

func (i IncomeController) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		param := vars["id"]
		id, err := strconv.Atoi(param)
		if err != nil {
			panic(err.Error())
		}
		income := model.IncomeModel{}.Find(id)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(income)
	}
}

func (i IncomeController) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var income model.IncomeModel
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &income)
		w.Header().Add("Content-Type", "application/json")
		iNew, err := income.Insert()
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(iNew)
		}
	}
}

func (i IncomeController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		income := model.IncomeModel{}
		vars := mux.Vars(r)
		param := vars["id"]
		id, _ := strconv.Atoi(param)
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &income)
		w.Header().Add("Content-Type", "application/json")
		iNew, err := income.Update(id)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(iNew)
		}
	}
}

func (i IncomeController) Destroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		income := model.IncomeModel{}
		vars := mux.Vars(r)
		param := vars["id"]
		id, _ := strconv.Atoi(param)
		w.Header().Add("Content-Type", "application/json")
		iNew, err := income.Destroy(id)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(iNew)
		}
	}
}
