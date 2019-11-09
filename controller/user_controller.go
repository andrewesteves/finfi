package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/andrewesteves/finfi/model"
	"github.com/gorilla/mux"
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

func (u UserController) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		param := vars["id"]
		id, err := strconv.Atoi(param)
		if err != nil {
			panic(err.Error())
		}
		user := model.UserModel{}.Find(id)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
