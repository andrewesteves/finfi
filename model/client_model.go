package model

type ClientModel struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}

func (c ClientModel) All() []ClientModel {
	return []ClientModel{
		ClientModel{
			Name:  "Bill Gates",
			Email: "bill@microsoft.com",
		},
		ClientModel{
			Name:  "Larry Pages",
			Email: "Larry@google.com",
		},
	}
}
