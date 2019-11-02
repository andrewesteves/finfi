package controller

import (
	"encoding/json"
	"net/http"

	"github.com/andrewesteves/finfi/model"
)

type IncomeController struct{}

func (i IncomeController) Routes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			incomes := model.IncomeModel{}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(incomes.All())
		}
	}
}
