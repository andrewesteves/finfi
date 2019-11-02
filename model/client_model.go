package model

type ClientModel struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}
