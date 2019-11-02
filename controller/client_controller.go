package controller

import (
	"encoding/json"
	"net/http"

	"github.com/andrewesteves/finfi/model"
)

type ClientController struct{}

func (c ClientController) Routes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			clients := model.ClientModel{}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(clients.All())
		}
	}
}
