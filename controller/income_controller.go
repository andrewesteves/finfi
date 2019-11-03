package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
		id := vars["id"]
		income := model.IncomeModel{
			Title: fmt.Sprintf("Item %v", id),
			Client: model.ClientModel{
				Name:  "Bill Gates",
				Email: "bill@microsoft.com",
			},
			Description:  "Lorem ipsum...",
			Status:       "Paid",
			Installments: 0,
			Total:        100.00,
		}
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
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(income)
	}
}

func (i IncomeController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		income := model.IncomeModel{
			Title: fmt.Sprintf("Item %v", id),
			Client: model.ClientModel{
				Name:  "Bill Gates",
				Email: "bill@microsoft.com",
			},
			Description:  "Lorem ipsum...",
			Status:       "Paid",
			Installments: 0,
			Total:        100.00,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(income)
	}
}

func (i IncomeController) Destroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		income := model.IncomeModel{
			Title: fmt.Sprintf("Item %v", id),
			Client: model.ClientModel{
				Name:  "Bill Gates",
				Email: "bill@microsoft.com",
			},
			Description:  "Lorem ipsum...",
			Status:       "Paid",
			Installments: 0,
			Total:        100.00,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(income)
	}
}
