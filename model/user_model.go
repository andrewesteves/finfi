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

func (u UserModel) Insert() (UserModel, []Errors) {
	if hasErrors := userValidate(u); len(hasErrors) > 0 {
		return u, hasErrors
	}

	db := storage.Connection()
	defer db.Close()
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		panic(err.Error())
	}
	u.Password = string(hash)
	rs, err := db.Exec("INSERT INTO users (name, email, password, role) VALUES (?, ?, ?, ?)", u.Name, u.Email, u.Password, u.Role)
	if err != nil {
		panic(err.Error())
	}
	id, err := rs.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	u.ID = int(id)
	return u, nil
}

func (u UserModel) Update(id int) (UserModel, []Errors) {
	if hasErrors := userValidate(u); len(hasErrors) > 0 {
		return u, hasErrors
	}

	db := storage.Connection()
	defer db.Close()
	user := u.Find(id)
	if u.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
		if err != nil {
			panic(err.Error())
		}
		u.Password = string(hash)
	} else {
		u.Password = user.Password
	}
	rs, err := db.Prepare("UPDATE users SET name = ?, email = ?, password = ?, role = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	rs.Exec(u.Name, u.Email, u.Password, u.Role, id)
	u.ID = id
	return u, nil
}

func (u UserModel) Destroy(id int) UserModel {
	db := storage.Connection()
	defer db.Close()
	rs, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	rs.Exec(id)
	u.ID = id
	return u
}

func userValidate(u UserModel) []Errors {
	var errs []Errors
	if u.Name == "" {
		errs = append(errs, Errors{"name", "The name field is required"})
	}
	if u.Email == "" {
		errs = append(errs, Errors{"email", "The e-mail field is required"})
	}
	return errs
}
