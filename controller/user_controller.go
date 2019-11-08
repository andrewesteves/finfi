package controller

import (
	"encoding/json"
	"net/http"

	"github.com/andrewesteves/finfi/model"
)

type UserController struct{}

func (u UserController) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := model.UserModel{}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users.All())
	}
}
