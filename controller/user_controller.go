package controller

import (
	"encoding/json"
	"io/ioutil"
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

func (u UserController) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.UserModel
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &user)
		w.Header().Add("Content-Type", "application/json")
		uNew, err := user.Insert()
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(uNew)
		}
	}
}

func (u UserController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.UserModel
		vars := mux.Vars(r)
		param := vars["id"]
		id, _ := strconv.Atoi(param)
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &user)
		w.Header().Add("Content-Type", "application/json")
		uNew, err := user.Update(id)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(uNew)
		}
	}
}

func (u UserController) Destroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.UserModel
		vars := mux.Vars(r)
		param := vars["id"]
		id, _ := strconv.Atoi(param)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user.Destroy(id))
	}
}
