package controller

import (
	"encoding/json"
	"net/http"

	"github.com/andrewesteves/finfi/model"
)

type UserController struct{}

func (u UserController) Routes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			users := model.UserModel{}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(users.All())
		}
	}
}
