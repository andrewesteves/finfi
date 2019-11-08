package model

import "github.com/andrewesteves/finfi/storage"

type UserModel struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
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
