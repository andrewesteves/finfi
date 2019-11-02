package model

type UserModel struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (u UserModel) All() []UserModel {
	return []UserModel{}
}
