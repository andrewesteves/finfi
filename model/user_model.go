package model

import (
	"github.com/andrewesteves/finfi/storage"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	ApiToken string `json:"api_token"`
}

func (u UserModel) All() []UserModel {
	var users []UserModel
	db := storage.Connection()
	defer db.Close()

	rs, err := db.Query("SELECT id, name, email, role FROM users")
	if err != nil {
		panic(err.Error())
	}

	for rs.Next() {
		var user UserModel
		err = rs.Scan(&user.ID, &user.Name, &user.Email, &user.Role)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	return users
}

func (u UserModel) Find(id int) UserModel {
	var user UserModel
	db := storage.Connection()
	defer db.Close()
	err := db.QueryRow("SELECT id, name, email, role FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Role)
	if err != nil {
		panic(err.Error())
	}
	return user
}

func (u UserModel) Insert() UserModel {
	db := storage.Connection()
	defer db.Close()
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		panic(err.Error())
	}
	u.Password = string(bytes)
	rs, err := db.Exec("INSERT INTO users (name, email, password, role) VALUES (?, ?, ?, ?)", u.Name, u.Email, u.Password, u.Role)
	if err != nil {
		panic(err.Error())
	}
	id, err := rs.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	u.ID = int(id)
	return u
}
